package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"main/app_interfaces"
	"main/entities"
	"main/infrastructure"
	"main/usecases"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL   = 12 * time.Hour
	saltCrypto = "35edtryuiojhgytfe3"
	signingKey = ",;mkljhgffxdgcfhvg"
)

type tokenClaims struct {
	jwt.Claims
	id int
}

type AuthService struct {
	usecases.Authorization
	app_interfaces.TokenAuth
	app_interfaces.UserInfo
	repos infrastructure.Authorization
}

func (serv *AuthService) AddUser(user entities.User) (int, error) {
	user.Password = generateHashedPassword(user.Password)
	return serv.repos.AddUser(user)
}

func (serv *AuthService) GetUser(fullName, password string) (*entities.User, error) {
	return nil, nil
}

func (serv *AuthService) GenerateToken(fullName, password string) (string, error) {
	user, err := serv.repos.GetUser(fullName, password)
	if err != nil {
		return "", err
	}
	fmt.Print(user)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, tokenClaims{jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	},user.Id})
	return token.SignedString([]byte(signingKey))
}

func (*AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method");
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims); if !ok {
		return 0, errors.New("token type is not of type *tokenClaims")
	}
	return claims.id, nil
}

func (serv *UserInfoService) GetRole(userId entities.UserRole) (entities.UserRole, error) {
	return serv.repos.GetRole(userId)
}

func NewAuthService(repos infrastructure.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func generateHashedPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltCrypto)))
}
