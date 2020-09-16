package controllers

import (
	"../config"
	"../models"
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	token, err := models.VerifyToken(r)
	if err != nil {
		return
	}
	if token.Valid {
		DeleteRefresh(token)
	}
}

func DeleteRefresh(token *jwt.Token) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"refreshtoken", token.Raw}}
	_, _ = config.GetDB().DeleteOne(ctx, filter)
}
