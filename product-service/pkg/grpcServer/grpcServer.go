package grpcServer

import (
	"PR10_1/product-service/pkg/api"
	"PR10_1/product-service/pkg/postgre"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type PrGRPCServer struct {
	api.UnimplementedProductServiceServer
	db postgre.Storage
}

func NewGRPCServer(db *gorm.DB) *PrGRPCServer {
	return &PrGRPCServer{
		db: postgre.NewProductServiceStorage(db),
	}
}

// InsertProduct ...
func (p *PrGRPCServer) InsertProduct(ctx context.Context, req *api.InsertProductRequest) (*api.InsertProductResponse, error) {
	product := postgre.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	fmt.Println("Hi")

	_, err := p.db.InsertProduct(ctx, product)
	if err != nil {
		return &api.InsertProductResponse{
			Result: &api.InsertProductResponse_IsSuccessful{
				IsSuccessful: fmt.Sprintf("%v", err),
			},
		}, err
	}

	return &api.InsertProductResponse{
		Result: &api.InsertProductResponse_Product{
			Product: &api.Product{
				Name:  product.Name,
				Price: product.Price,
			},
		},
	}, nil
}

// GetAllProduct ...
func (p *PrGRPCServer) GetAllProduct(ctx context.Context, req *api.GetAllProductRequest) (*api.GetAllProductResponse, error) {
	//TODO implement me
	panic("implement me")

}

// GetProduct ...
func (p *PrGRPCServer) GetProduct(ctx context.Context, req *api.GetProductRequest) (*api.GetProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteProduct ...
func (p *PrGRPCServer) DeleteProduct(ctx context.Context, req *api.DeleteProductRequest) (*api.DeleteProductResponse, error) {
	return &api.DeleteProductResponse{
		IsSuccessful: false,
	}, nil
}
