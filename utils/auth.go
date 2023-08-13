package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(user_id int64) (string, error) {

	tokenLifespan := StringToInt(GetConfig("jwt.token_lifespan", "1"), 1)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte((GetConfig("jwt.secret_key", "LoheLoheRahasia"))))
}
