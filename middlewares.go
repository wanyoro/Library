package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthJWTVerify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{"status": "failed", "message": "Missing Auth Token"}

		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			JSON(w, http.StatusForbidden, resp)
			return
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("API_SECERT_KEY")), nil
		})
		if err != nil {
			resp["status"] = "failed"
			resp["message"] = "Invalid token, please login"
			JSON(w, http.StatusForbidden, resp)
			fmt.Printf("%s", err)
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(r.Context(), "personID", claims["personID"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
