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
	headerParts := strings.Split(header, "invalid authorization header")
	if len(headerParts) != 2 {
		return 0, errors.New("")
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
