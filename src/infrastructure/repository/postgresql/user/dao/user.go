package dao

type User struct {
	Id       int64 `gorm:"primaryKey"`
	Name     string
	UserName string `gorm:"uniqueIndex"`
	Email    string
	Password string
	Rol      string
}
