package main

import (
	"log"
	"net"
	"task-service/internal/db"
	"task-service/internal/handler"
	"task-service/pb/task_pb"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	grpcServer := grpc.NewServer()
	task_pb.RegisterTaskServiceServer(grpcServer, handler.NewTaskHandler())

	log.Println("gRPC сервер запущен на порту 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}
}
