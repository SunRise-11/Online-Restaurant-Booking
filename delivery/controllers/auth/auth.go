package auth

import (
	"Restobook/delivery/common"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateTokenAuth(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.JWT_SECRET_KEY))
}

func CreateTokenAuthRestaurant(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["restoid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.JWT_SECRET_KEY))
}
