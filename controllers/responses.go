package controllers

import (
	"encoding/json"
	"log"
	domainErrors "main/domain/entities/domain_errors"
	serviceErrors "main/service/errors"
	"net/http"
)


func newErrorResponse(writer http.ResponseWriter, code int, message string) {
	log.Println(message)
	writer.WriteHeader(code)
	res, _ := json.Marshal(message)
	writer.Write(res)
}

func okResponse(writer http.ResponseWriter) {
	res, _ := json.Marshal("ok")
	writer.Write(res)
}

func lastErrorHandling(writer http.ResponseWriter, err error) {
	switch err.(type) {
	case *serviceErrors.RoleError:
		newErrorResponse(writer, http.StatusUnauthorized, err.Error())
		return
	case *domainErrors.InvalidField:
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	default:
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}