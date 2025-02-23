package service

import (
	"errors"
	"main/domain/entities"
	"main/service/entities_models/request"
	"main/service/entities_models/response"
	request_mappers "main/service/mappers/request"
	response_mappers "main/service/mappers/response"
	"main/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: DTOS
const (
	tokenTTL   = 12 * time.Hour
	signingKey = ",;mkljhgffxdgcfhvg"
)

type tokenClaims struct {
	jwt.Claims
	id int
}

type roleAccess interface {
	GetUserRole(userId int) (entities.UserRole, error)
}
type TokenAuth interface {
	GenerateToken(fullName, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	roleAccess
}
type Authorization interface {
	AddUser(user request.ClientSignUpModel) error
	// AddAdmin(admin request.AdminSignUpModel) error
	// AddManager(manager request.ManagerSignUpModel) error 
	// AddOuterSpecialist(manager request.ManagerSignUpModel) error
	GetUser(username, password string) (*response.UserAuthModel, error)
}

type AuthService struct {
	Authorization
	TokenAuth
	//RoleAccess
	repos AuthorizationRepository
}

func (serv *AuthService) AddUser(user request.ClientSignUpModel) error {
	user.Password = utils.GenerateHashedPassword(user.Password)
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

func (serv *AuthService) GenerateToken(fullName, password string) (string, error) {
	user, err := serv.repos.GetUser(fullName, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, tokenClaims{jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	},
		user.Id(),
	})
	return token.SignedString([]byte(signingKey))
}

func (serv *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	return claims.id, nil
}

func NewAuthService(repos AuthorizationRepository) *AuthService {
	return &AuthService{
		repos: repos,
	}
}
