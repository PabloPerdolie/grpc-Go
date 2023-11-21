package grpcServer

import (
	"PR10_1/product-service/pkg/api"
	"context"
)

type Server interface {
	InsertProduct(ctx context.Context, req *api.InsertProductRequest) (*api.InsertProductResponse, error)
	GetAllProduct(ctx context.Context, req *api.GetAllProductRequest) (*api.GetAllProductResponse, error)
	GetProduct(ctx context.Context, req *api.GetProductRequest) (*api.GetProductResponse, error)
	DeleteProduct(ctx context.Context, req *api.DeleteProductRequest) (*api.DeleteProductResponse, error)
}
