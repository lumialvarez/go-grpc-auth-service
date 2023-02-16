package login

import (
	"context"
	"fmt"
	"github.com/lumialvarez/go-common-tools/hash"
	domainError "github.com/lumialvarez/go-grpc-auth-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
	"time"
)

type Repository interface {
	GetByEmail(email string) (*user.User, error)
	GetByUserName(username string) (*user.User, error)
	Save(user *user.User) (*user.User, error)
}

type JwtServiceUser interface {
	GenerateToken(user *user.User) (signedToken string, err error)
	ValidateToken(signedToken string) (*user.User, error)
}

type NotificationService interface {
	PublishNotification(title string, detail string) error
}

type UseCaseLoginUser struct {
	repository          Repository
	jwtService          JwtServiceUser
	notificationService NotificationService
}

func NewUseCaseLoginUser(repository Repository, jwtService JwtServiceUser, notificationService NotificationService) UseCaseLoginUser {
	return UseCaseLoginUser{repository: repository, jwtService: jwtService, notificationService: notificationService}
}

func (uc UseCaseLoginUser) Execute(ctx context.Context, domainUser *user.User) (*user.User, error) {
	dbUser, err := uc.repository.GetByUserName(domainUser.UserName())
	if err != nil {
		uc.processFailAttempt(domainUser, "Invalid credentials")
		return nil, domainError.NewInvalidCredentials("Invalid credentials")
	}

	match := hash.CheckPasswordHash(domainUser.Password(), dbUser.Password())

	if !match {
		uc.processFailAttempt(domainUser, "Invalid credentials")
		return nil, domainError.NewInvalidCredentials("Invalid credentials")
	}

	if !dbUser.Status() {
		uc.processFailAttempt(domainUser, "User Inactive")
		return nil, domainError.NewInactive("User Inactive")
	}

	token, _ := uc.jwtService.GenerateToken(dbUser)
	dbUser.SetToken(token)

	return dbUser, nil
}

func (uc UseCaseLoginUser) processFailAttempt(domainUser *user.User, reason string) {
	title := "Login attempt failed"
	detail := fmt.Sprintf("A failed attempt is registed in %s with this credentials User: %s Password: %s, reason: %s", time.Now().String(), domainUser.UserName(), domainUser.Password(), reason)
	uc.notificationService.PublishNotification(title, detail)
}
