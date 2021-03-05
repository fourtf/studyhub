package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserID         uint
	Name           string
	Email          string
	StandardClaims *jwt.StandardClaims
}
