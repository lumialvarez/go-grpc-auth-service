package errorGrpc

import (
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type APIResponseProvider struct {
}

func NewAPIResponseProvider() *APIResponseProvider {
	return &APIResponseProvider{}
}

func (a *APIResponseProvider) ToAPIResponse(err error) error {
	switch err.(type) {
	case domainError.AlreadyExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case domainError.NotFound:
		return status.Error(codes.NotFound, err.Error())
	case domainError.InvalidCredentials:
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
