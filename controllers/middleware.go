package controllers

import (
	"errors"
	"main/domain/entities"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

func (controller *Controller) userIdentity(req *http.Request) (int, error) {
	header := req.Header.Get(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty header")
	}
	headerParts := strings.Split(header, "invalid authorization header")
	if len(headerParts) != 2 {
		return 0, errors.New("")
	}
	userId, err := controller.services.ParseToken(headerParts[1])
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (controller *Controller) userRole(req *http.Request) (entities.UserRole, error) {
	var userId int
	var err error
	header := req.Header.Get(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, errors.New("invalid authorization header")
	}
	userId, err = controller.services.ParseToken(headerParts[1])
	if err != nil {
		return 0, err
	}
	role, err := controller.services.GetUserRole(userId)
	if err != nil {
		return 0, err
	}
	return role, nil
}
