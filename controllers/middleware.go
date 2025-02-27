package controllers

import (
	"errors"
	"main/domain/entities"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

func (middleware *Middleware) userIdentity(req *http.Request) (int, error) {
	header := req.Header.Get(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, errors.New("invalid header length")
	}
	userId, err := middleware.authMiddleware.ParseToken(headerParts[1])
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (middleware *Middleware) userRole(req *http.Request) (entities.UserRole, error) {
	userId, err := middleware.userIdentity(req)
	if err != nil {
		return 0,err
	}
	role, err :=  middleware.authMiddleware.GetUserRole(userId)
	if err != nil {
		return 0, err
	}
	return role, nil
}

func (middleware *Middleware) enableCors(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
    writer.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
