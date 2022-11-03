package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/service/jwt/user/dto"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomain(jwtClaims *dto.JwtClaims) *user.User {
	return user.NewUser(jwtClaims.Id, jwtClaims.Email, "", "")
}
