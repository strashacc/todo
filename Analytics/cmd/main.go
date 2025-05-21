package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // импорт библиотеки

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"analytics/config"
	grpcDelivery "analytics/internal/delivery/grpc"
	pb "analytics/internal/delivery/grpc/pb"
	"analytics/internal/delivery/rest"
	mongoRepo "analytics/internal/repository/mongo"
	"analytics/internal/usecase"
)

func main() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(ctx)
	db := client.Database(cfg.DatabaseName)

	repo := mongoRepo.NewTaskRepository(db)
	uc := usecase.NewTaskUsecase(repo)

	// Запуск gRPC сервера
	go func() {
		lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
		if err != nil {
			log.Fatalf("gRPC listen error: %v", err)
		}

		s := grpcDelivery.NewTaskHandler(uc)
		grpcServer := grpcLib.NewServer()
		pb.RegisterTaskServiceServer(grpcServer, s)
		reflection.Register(grpcServer)

		log.Println("gRPC server listening on port", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC serve error: %v", err)
		}
	}()

	// Запуск REST сервера
	r := gin.Default()
	rest.NewRESTHandler(r, uc)

	go func() {
		log.Println("REST server listening on port", cfg.RESTPort)
		if err := r.Run(":" + cfg.RESTPort); err != nil {
			log.Fatalf("REST server error: %v", err)
		}
	}()

	// Graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down...")
}
