package service

import (
	"errors"
	"main/domain/entities"
	"main/service/entities_models/request"
	"main/service/entities_models/response"
	request_mappers "main/service/mappers/request"
	response_mappers "main/service/mappers/response"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
	"main/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = ",;mkljhgffxdgcfhvg"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Id int `json:"user_id"`
}

type AuthService struct {
	serviceInterfaces.Authorization
	serviceInterfaces.TokenAuth
	//RoleAccess
	repos repository.AuthorizationRepository
}

func (serv *AuthService) AddUser(user request.ClientSignUpModel) error {
	userEntity, err := request_mappers.ToUserEntitiy(user, serv.repos)
	if err != nil {
		return err
	}
	return serv.repos.AddUser(*userEntity)
}

func (serv *AuthService) GetUser(fullName, password string) (*response.UserAuthModel, error) {
	usrEntity, err := serv.repos.GetUser(fullName, password)
	if err != nil {
		return nil, err
	}
	return response_mappers.ToUserAuthModel(*usrEntity), err
}

func (serv *AuthService) GetUserRole(userId int) (entities.UserRole, error) {
	return serv.repos.GetUserRole(userId)
}

func (serv *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := serv.repos.GetUser(email, utils.GenerateHashedPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	},
		user.Id(),
	})
	return token.SignedString([]byte(signingKey))
}

func (serv *AuthService) ParseToken(accessToken string) (int, error) {
	customClaims := &tokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token type is not of type *tokenClaims")
	}
	return claims.Id, nil
}

func NewAuthService(repos repository.AuthorizationRepository) *AuthService {
	return &AuthService{
		repos: repos,
	}
}
