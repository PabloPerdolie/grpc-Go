package postgre

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID string `gorm:"uniqueIndex;autoIncrement"`
	Name      string
	Price     int64
}
