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

func (pr *productServiceStorage) GetProduct(ctx context.Context, id string) (product Product, err error) {
	result := pr.db.Where("ID = ?", id).Find(&product)
	if result.Error != nil {
		log.Printf("failed to find product, %v", result.Error.Error())
		return product, result.Error
	}
	return product, err
}

func (pr *productServiceStorage) GetAllProduct(ctx context.Context) (products []Product, err error) {
	result := pr.db.Find(&products)
	if result.Error != nil {
		log.Printf("failed to find products, %v", result.Error.Error())
		return products, result.Error
	}
	return products, err
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
	var product Product
	result := pr.db.Where("ID = ?", id).Find(&product)
	if result.Error != nil {
		log.Printf("failed to delete product, %v", result.Error.Error())
		return result.Error
	}
	return nil
}
