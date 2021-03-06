package models

import "github.com/dgrijalva/jwt-go"

//Token that is used to generate the token string the user needs for authed paths
type Token struct {
	UserID         uint
	Name           string
	Email          string
	StandardClaims *jwt.StandardClaims
}
