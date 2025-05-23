package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"collaboration/config"
	"collaboration/db/mongodb"
	grpcServer "collaboration/internal/server/grpc"
	pb "collaboration/pb/team_pb"

	dotenv "github.com/joho/godotenv"
	grpcLib "google.golang.org/grpc"
)

func main() {
	// Set log output to file
	file, err := os.Create("/logs/team.log")
	if err != nil {
		log.Println("Failed setting log output to file")
	} else {
		defer file.Close()
		log.SetOutput(file)
	}

	// Load config
	if err := dotenv.Load(); err != nil {
		log.Println("File .env was not found, continuing with default values")
	}
	cfg := config.Load()

	if err := mongodb.StartConnection(cfg.MONGO_URI, cfg.DB); err != nil {
		log.Fatalf("Couldn't connect to database: %s\n", err.Error())
	}

	// Start grpc server
	go func () {
		lis, err := net.Listen("tcp", ":" + cfg.GRPC_PORT)
		if err != nil {
			log.Fatalf("gRPC listen error: %v", err)
		}

		s := grpcServer.NewServer()
		grpcServer := grpcLib.NewServer()
		pb.RegisterTeamServiceServer(grpcServer, s)

		log.Println("gRPC server listening on port", cfg.GRPC_PORT)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC serve error: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Shutting down...")
}