package models

import "github.com/dgrijalva/jwt-go"

//Claims that is used to generate the token string the user needs for authed paths
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

type key int

//UserKey is a unique key
const UserKey = key(1)
