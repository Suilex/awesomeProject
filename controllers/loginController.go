package controllers

import (
	"../config"
	"../models"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("id")

	if userId != "" {
		CreateToken(w, userId)
	} else {
		_, _ = w.Write([]byte("error"))
	}
}

func CreateToken(w http.ResponseWriter, userId string) {
	Aclaims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS512"), Aclaims)
	access, _ := accessToken.SignedString([]byte(models.AccessSecret))

	Rclaims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS512"), Rclaims)
	refresh, _ := refreshToken.SignedString([]byte(models.RefreshPassword))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	table := models.TableReq{
		UserId:       userId,
		RefreshToken: refresh,
	}
	_, _ = config.GetDB().InsertOne(ctx, table)
	_, _ = w.Write([]byte("AccessToken  " + access))
	_, _ = w.Write([]byte("\n"))
	_, _ = w.Write([]byte("RefreshToken  " + refresh))
}
