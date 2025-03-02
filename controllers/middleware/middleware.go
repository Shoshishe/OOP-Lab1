package middleware

import (
	"errors"
	"main/domain/entities"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

type Middleware struct {
	authMiddleware serviceInterfaces.TokenAuth
}

func NewMiddleware(authMiddleware serviceInterfaces.TokenAuth) *Middleware {
	return &Middleware{
		authMiddleware: authMiddleware,
	}
}

func (middleware *Middleware) UserIdentity(req *http.Request) (int, error) {
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

func (middleware *Middleware) UserRole(req *http.Request) (entities.UserRole, error) {
	userId, err := middleware.UserIdentity(req)
	if err != nil {
		return 0,err
	}
	role, err :=  middleware.authMiddleware.GetUserRole(userId)
	if err != nil {
		return 0, err
	}
	return role, nil
}

func (middleware *Middleware) EnableCors(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
    writer.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
