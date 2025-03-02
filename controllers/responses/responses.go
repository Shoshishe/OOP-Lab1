package controllerResponse

import (
	"encoding/json"
	"log"
	domainErrors "main/domain/entities/domain_errors"
	serviceErrors "main/service/errors"
	"net/http"
)


func NewErrorResponse(writer http.ResponseWriter, code int, message string) {
	log.Println(message)
	writer.WriteHeader(code)
	res, _ := json.Marshal(message)
	writer.Write(res)
}

func OkResponse(writer http.ResponseWriter) {
	res, _ := json.Marshal("ok")
	writer.Write(res)
}

func LastErrorHandling(writer http.ResponseWriter, err error) {
	switch err.(type) {
	case *serviceErrors.RoleError:
		NewErrorResponse(writer, http.StatusUnauthorized, err.Error())
		return
	case *domainErrors.InvalidField:
		NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	default:
		NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}