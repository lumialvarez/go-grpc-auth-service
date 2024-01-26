package dto

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	jwt.RegisteredClaims
	Id       int64
	UserName string
	Rol      string
}
