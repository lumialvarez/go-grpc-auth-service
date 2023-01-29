package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user/dao"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomain(daoUser *dao.User) *user.User {
	var domainNotifications []user.Notification
	for _, daoNotification := range daoUser.UserNotification {
		domainNotification := user.NewNotification(daoNotification.Id, daoNotification.Title, daoNotification.Detail, daoNotification.Date, daoNotification.Read)
		domainNotifications = append(domainNotifications, *domainNotification)
	}

	return user.NewUser(daoUser.Id, daoUser.Name, daoUser.UserName, daoUser.Email, daoUser.Password, "", user.Role(daoUser.Rol), daoUser.Status, domainNotifications)
}

func (m Mapper) ToDAO(domainUser *user.User) *dao.User {
	daoUser := dao.User{
		Id:       domainUser.Id(),
		Name:     domainUser.Name(),
		UserName: domainUser.UserName(),
		Email:    domainUser.Email(),
		Password: domainUser.Password(),
		Rol:      string(domainUser.Role()),
		Status:   domainUser.Status(),
	}
	return &daoUser
}
