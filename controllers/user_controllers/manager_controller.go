package userControllers

import (
	"main/controllers/middleware"
	controllerResponse "main/controllers/responses"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type managerController struct {
	serv    serviceInterfaces.ManagerService
	middleware middleware.Middleware
}

func (controller *managerController) ApproveCredit(writer http.ResponseWriter, req *http.Request) {
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
	requestIdArg := req.PathValue("request_id")
	requestId, err := strconv.Atoi(requestIdArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "invalid path arg")
		return
	}
	err = controller.serv.ApproveCredit(requestId, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func NewMangerController(serv serviceInterfaces.ManagerService, middleware middleware.Middleware) *managerController {
	return &managerController{
		serv: serv, 
		middleware: middleware,
	}
}