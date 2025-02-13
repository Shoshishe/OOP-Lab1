package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)


type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(writer http.ResponseWriter, code int, message string) {
	log.Println(message)
	writer.WriteHeader(code)
	res, _ := json.Marshal(message)
	writer.Write(res)
}