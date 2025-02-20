package controllers

import (
	"encoding/json"
	"log"
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