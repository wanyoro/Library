package main

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//create auth token

func AuthToken(uid uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["PersonID"] = uid
	claims["IssuedAt"] = time.Now().Unix()
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(os.Getenv("API_SECERT_KEY")))
}
