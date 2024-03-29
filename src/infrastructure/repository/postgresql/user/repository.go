package repositoryUser

import (
	"github.com/lumialvarez/go-common-tools/platform/postgresql"
	"github.com/lumialvarez/go-grpc-auth-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user/dao"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/repository/postgresql/user/mapper"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Repository struct {
	postgresql postgresql.Client
	mapper     mapper.Mapper
}

func Init(config config.Config) Repository {
	postgresqlClient := postgresql.Init(config.DBUrl)
	postgresqlClient.DB.AutoMigrate(dao.User{})
	postgresqlClient.DB.AutoMigrate(dao.UserNotification{})
	return Repository{postgresql: postgresqlClient, mapper: mapper.Mapper{}}
}

func (repository *Repository) GetById(id int64) (*user.User, error) {
	var daoUser dao.User
	result := repository.postgresql.DB.Where(&dao.User{Id: id}).Preload("UserNotification").First(&daoUser)
	if result.Error != nil {
		return nil, result.Error
	}
	domainUser := repository.mapper.ToDomain(&daoUser)

	return domainUser, nil
}

func (repository *Repository) GetByEmail(email string) (*user.User, error) {
	var daoUser dao.User
	result := repository.postgresql.DB.Where(&dao.User{Email: email}).First(&daoUser)
	if result.Error != nil {
		return nil, result.Error
	}
	domainUser := repository.mapper.ToDomain(&daoUser)

	return domainUser, nil
}

func (repository *Repository) GetByUserName(username string) (*user.User, error) {
	var daoUser dao.User
	result := repository.postgresql.DB.Where(&dao.User{UserName: username}).First(&daoUser)
	if result.Error != nil {
		return nil, result.Error
	}
	domainUser := repository.mapper.ToDomain(&daoUser)

	return domainUser, nil
}

func (repository *Repository) Save(domainUser *user.User) (*user.User, error) {
	daoUser := repository.mapper.ToDAO(domainUser)
	result := repository.postgresql.DB.Create(&daoUser)
	if result.Error != nil {
		return nil, result.Error
	}
	domainUser = repository.mapper.ToDomain(daoUser)

	return domainUser, nil
}

func (repository *Repository) Update(domainUser *user.User) (*user.User, error) {
	daoUser := repository.mapper.ToDAO(domainUser)
	updateUser := make(map[string]interface{})
	updateUser["name"] = daoUser.Name
	updateUser["rol"] = daoUser.Rol
	updateUser["email"] = daoUser.Email
	updateUser["status"] = daoUser.Status
	if len(daoUser.Password) > 0 {
		updateUser["password"] = daoUser.Password
	}

	result := repository.postgresql.DB.Model(&daoUser).Updates(updateUser)
	if result.Error != nil {
		return nil, result.Error
	}
	domainUser = repository.mapper.ToDomain(daoUser)

	return domainUser, nil
}

func (repository *Repository) GetAll() (*[]user.User, error) {
	var daoUsers []dao.User
	result := repository.postgresql.DB.Find(&daoUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	var domainUsers []user.User
	for _, daoUser := range daoUsers {
		domainUsers = append(domainUsers, *repository.mapper.ToDomain(&daoUser))
	}

	return &domainUsers, nil
}

func (repository *Repository) CreateNotificationToAdminUsers(notification user.Notification) error {
	var daoUsers []dao.User
	result := repository.postgresql.DB.Where(&dao.User{Rol: user.RolAdmin}).Preload("UserNotification").Find(&daoUsers)
	if result.Error != nil {
		return result.Error
	}

	for _, daoUser := range daoUsers {
		daoNotification := repository.mapper.ToDAONotification(daoUser.Id, &notification)
		result := repository.postgresql.DB.Where(&daoUser).Model(&daoNotification).Create(&daoNotification)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (repository *Repository) MarkReadNotification(userId int64, notificationId int64) error {
	var daoUser dao.User
	updateUser := make(map[string]interface{})
	updateUser["read"] = "true"

	result := repository.postgresql.DB.Where(&dao.User{Id: userId}).Preload("UserNotification").First(&daoUser)
	if result.Error != nil {
		return result.Error
	}

	for _, notification := range daoUser.UserNotification {
		if notification.Id == notificationId {
			result := repository.postgresql.DB.Model(&notification).Updates(updateUser)
			if result.Error != nil {
				return result.Error
			}
			break
		}
	}

	return nil
}
