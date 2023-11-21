package postgre

import (
	"context"
	"gorm.io/gorm"
	"log"
)

type productServiceStorage struct {
	context context.Context
	db      *gorm.DB
}

func NewProductServiceStorage(db *gorm.DB) Storage {
	return &productServiceStorage{
		context: context.Background(),
		db:      db,
	}
}

func (pr *productServiceStorage) GetProduct(ctx context.Context, id string) (Product, error) {
	// TODO implement me
	panic("implement me")
}

func (pr *productServiceStorage) GetAllProduct(ctx context.Context) ([]Product, error) {
	// TODO implement me
	panic("implement me")
}

func (pr *productServiceStorage) InsertProduct(ctx context.Context, prod Product) (string, error) {
	result := pr.db.Create(&prod)
	if result.Error != nil {
		log.Printf("failed to insert product, %v", result.Error.Error())
		return "", result.Error
	}
	return prod.ProductID, nil
}

func (pr *productServiceStorage) DeleteProduct(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}
