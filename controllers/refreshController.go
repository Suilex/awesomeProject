package controllers

import (
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	token, err := models.VerifyToken(r)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := fmt.Sprintf("%s", claims["userId"])

		DeleteRefresh(token)
		CreateToken(w, userId)
	}
}
