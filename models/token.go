package models

import "github.com/dgrijalva/jwt-go"


type Token struct {
	UserId int
	Username string
	jwt.StandardClaims
}