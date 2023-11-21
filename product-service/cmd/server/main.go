package main

import (
	"PR10_1/product-service/pkg/api"
	"PR10_1/product-service/pkg/grpcServer"
	"PR10_1/product-service/pkg/postgre"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

func main() {

	if err := godotenv.Load("C:\\Study\\3year\\5sem\\РСЧИР\\PR10_1\\product-service\\.env"); err != nil {
		log.Println("No .env file found")
		panic(err)
	}
	// TODO change settings for DB
	dsn := os.Getenv("PSQL_AUTH")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Автомиграции для создания таблицы Product
	err = db.AutoMigrate(&postgre.Product{})
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterProductServiceServer(server, grpcServer.NewGRPCServer(db))

	log.Println("Product Service is running on :50051")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
