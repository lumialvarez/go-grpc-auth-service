package dto

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	Id       int64
	UserName string
	Rol      string
}
