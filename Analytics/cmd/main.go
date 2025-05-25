package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"analytics/config"
	grpcDelivery "analytics/internal/delivery/grpc"
	pb "analytics/internal/delivery/grpc/pb"
	mongoRepo "analytics/internal/repository/mongo"
	"analytics/internal/usecase"
)

func main() {
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

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpcLib.NewServer()
	pb.RegisterAnalyticsServiceServer(grpcServer, grpcDelivery.NewAnalyticsHandler(uc))

	reflection.Register(grpcServer)

	log.Println("gRPC server listening on port", cfg.GRPCPort)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Graceful shutdown on SIGINT, SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
}
