package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	productpb "PR10_1/product-service/pkg/api"
	userpb "PR10_1/user-service/pkg/api"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Регистрация ProductService
	productEndpoint := "localhost:50051"
	productConn, err := grpc.DialContext(ctx, productEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial product service: %v", err)
	}
	defer productConn.Close()

	err = productpb.RegisterProductServiceHandler(ctx, mux, productConn)
	if err != nil {
		log.Fatalf("Failed to register product service handler: %v", err)
	}

	// Регистрация UserService
	userEndpoint := "localhost:50052"
	userConn, err := grpc.DialContext(ctx, userEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial product service: %v", err)
	}
	defer userConn.Close()

	err = userpb.RegisterUserServiceHandler(ctx, mux, userConn)
	if err != nil {
		log.Fatalf("Failed to register user service handler: %v", err)
	}
	// Запуск API Gateway
	gatewayAddr := ":8080"
	log.Printf("Starting API Gateway on %s", gatewayAddr)
	err = http.ListenAndServe(gatewayAddr, mux)
	if err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}

}
