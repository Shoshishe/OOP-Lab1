package userControllers

import (
	"encoding/json"
	"main/controllers/middleware"
	controllerResponse "main/controllers/responses"
	"main/service/entities_models/request"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type outerSpecialistController struct {
	serv       serviceInterfaces.OuterSpecialistService
	middleware middleware.Middleware
}

func (controller *outerSpecialistController) SendInfo(writer http.ResponseWriter, req *http.Request) {
	var paymentReqModel request.PaymentRequestModel
	usrId, err := controller.middleware.UserIdentity(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = json.NewDecoder(req.Body).Decode(&paymentReqModel)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	requestIdArg := req.PathValue("request_id")
	requestId, err := strconv.Atoi(requestIdArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.serv.SendInfoForPayment(requestId, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *outerSpecialistController) TransferRequest(writer http.ResponseWriter, req *http.Request) {
	var input request.TransferModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrId, err := controller.middleware.UserIdentity(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	input.TransferOwnerId = usrId
	err = controller.serv.TransferRequest(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func NewOuterSpecialistController(serv serviceInterfaces.OuterSpecialistService, middleware middleware.Middleware) *outerSpecialistController {
	return &outerSpecialistController{
		serv:       serv,
		middleware: middleware,
	}
}
