package postgre

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type userServiceStorage struct {
	context context.Context
	db      *gorm.DB
}

func NewUserServiceStorage(db *gorm.DB) Storage {
	return &userServiceStorage{
		context: context.Background(),
		db:      db,
	}
}

func (pr *userServiceStorage) InsertUser(ctx context.Context, usr User) (string, error) {
	result := pr.db.Create(&usr)
	if result.Error != nil {
		log.Printf("failed to insert user, %v", result.Error.Error())
		return "", result.Error
	}
	return usr.UserID, nil
}

func (pr *userServiceStorage) DeleteUser(ctx context.Context, id int32) error {
	var user User
	result := pr.db.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		log.Printf("failed to delete user, %v", result.Error.Error())
		return result.Error
	}
	result = pr.db.Delete(&user)
	if result.Error != nil {
		fmt.Printf("failed to delete user: " + result.Error.Error())
	} else if result.RowsAffected == 0 {
		fmt.Printf("no user record was deleted")
	} else {
		fmt.Println("User record deleted successfully")
	}
	return nil
}
