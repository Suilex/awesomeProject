package controllers

import (
	"../config"
	"../models"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func DelAll(w http.ResponseWriter, r *http.Request) {
	token, err := models.VerifyToken(r)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		userId := fmt.Sprintf("%s", claims["userId"])
		DeleteUserId(userId)
	}
}

func DeleteUserId(userId string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"userid", userId}}
	_, _ = config.GetDB().DeleteMany(ctx, filter)
}
