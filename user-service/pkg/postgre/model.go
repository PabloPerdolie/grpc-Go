package postgre

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"uniqueIndex;autoIncrement"`
	Username string
	Password string
	Name     string
}
