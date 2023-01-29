package dao

import (
	"time"
)

type User struct {
	//gorm.Model
	Id               int64 `gorm:"primaryKey"`
	Name             string
	UserName         string `gorm:"uniqueIndex"`
	Email            string
	Password         string
	Rol              string
	Status           bool
	UserNotification []UserNotification `gorm:"foreignKey:IdUser;references:Id"`
}

type UserNotification struct {
	Id     int64 `gorm:"primaryKey"`
	IdUser int64
	Title  string
	Detail string
	Date   time.Time `gorm:"default:CURRENT_TIMESTAMP",json:"date"`
	Read   bool      `gorm:"default:false"`
}
