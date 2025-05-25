package main

import (
	"log"
	"net"
	"net/http"

	"github.com/strashacc/todo/Notification/config"
	"github.com/strashacc/todo/Notification/handlers"
	pb "github.com/strashacc/todo/Notification/proto"
	"github.com/strashacc/todo/Notification/storage"
	"google.golang.org/grpc"
)

func main() {
	config.LoadEnv()
	db, err := storage.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}

	// gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, &handlers.NotificationServer{DB: db})
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("gRPC Notification Service started on :50051")
		grpcServer.Serve(lis)
	}()

	// REST API
	http.HandleFunc("/api/notifications", handlers.UserNotificationsREST(db))
	log.Println("REST API started on :8082")
	http.ListenAndServe(":8082", nil)
}
