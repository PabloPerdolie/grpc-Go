package main

import (
	"PR10_1/user-service/pkg/api"
	"PR10_1/user-service/pkg/grpcServer"
	"PR10_1/user-service/pkg/postgre"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

func main() {
	if err := godotenv.Load("C:\\Study\\3year\\5sem\\РСЧИР\\PR10_1\\user-service\\.env"); err != nil {
		log.Println("No .env file found")
		panic(err)
	}
	dsn := os.Getenv("PSQL_AUTH")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Автомиграции для создания таблицы User
	err = db.AutoMigrate(&postgre.User{})
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterUserServiceServer(server, grpcServer.NewGRPCServer(db))

	log.Println("User Service is running on :50052")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
