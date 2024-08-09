package infrastructure

import (
	"errors"
	domain "task_manager/Domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {}

type JWT interface {
	GenerateToken(loggedUser domain.User)(string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}


func NewJWtService()*JWTService{
	return &JWTService{}
}

var jwtSecret = []byte("secret")

func (Jwt *JWTService)GenerateToken(loggedUser domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  loggedUser.ID,
		"username": loggedUser.UserName,
		"role":     loggedUser.Role,
		"expires":  (time.Now().Add(5 * time.Minute)).Unix(),
	})
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("user not logged in")
	}
	return jwtToken, nil
}

func (Jwt *JWTService)ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
