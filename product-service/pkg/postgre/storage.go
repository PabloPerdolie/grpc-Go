package postgre

import (
	"context"
)

type Storage interface {
	GetProduct(ctx context.Context, id int32) (Product, error)
	GetAllProduct(ctx context.Context) ([]Product, error)
	InsertProduct(ctx context.Context, prod Product) (string, error)
	DeleteProduct(ctx context.Context, id int32) error
}
