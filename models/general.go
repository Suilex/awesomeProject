package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)


type Claims struct {
	UserId     string `json:"userId"`
	jwt.StandardClaims
}

type TableReq struct {
	UserId       string `json:"userId"`
	RefreshToken string `json:"refresh_token"`
}

var AccessSecret = "access_password"
var RefreshPassword = "refresh_password"

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(RefreshPassword), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}