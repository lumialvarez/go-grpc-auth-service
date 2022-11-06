package serviceJwtUser

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/service/jwt/user/dto"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/service/jwt/user/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"time"
)

type Service struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
	mapper.Mapper
}

func (s *Service) GenerateToken(user *user.User) (signedToken string, err error) {
	claims := &dto.JwtClaims{
		Id:    user.Id(),
		Email: user.Email(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.ExpirationHours)).Unix(),
			Issuer:    s.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(s.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *Service) ValidateToken(signedToken string) (*user.User, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*dto.JwtClaims)

	if !ok {
		//Fixme
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		//Fixme
		return nil, errors.New("JWT is expired")
	}
	//Fixme
	return s.ToDomain(claims), nil

}
