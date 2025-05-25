package handlers

import (
	"context"
	"time"

	"github.com/strashacc/todo/Notification/internal"
	"github.com/strashacc/todo/Notification/models"
	pb "github.com/strashacc/todo/Notification/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	DB *mongo.Database
}

func (s *NotificationServer) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.NotificationResponse, error) {
	notif := models.Notification{
		Type:      models.NotificationType(req.Type),
		Message:   req.Message,
		Read:      false,
		CreatedAt: time.Now(),
	}
	if req.UserId != "" {
		id, _ := primitive.ObjectIDFromHex(req.UserId)
		notif.UserID = &id
	}
	if req.TeamId != "" {
		id, _ := primitive.ObjectIDFromHex(req.TeamId)
		notif.TeamID = &id
	}
	if req.TaskId != "" {
		id, _ := primitive.ObjectIDFromHex(req.TaskId)
		notif.TaskID = id
	}
	res, err := s.DB.Collection("notifications").InsertOne(ctx, notif)
	if err != nil {
		return nil, err
	}
	notif.ID = res.InsertedID.(primitive.ObjectID)
	go internal.PushNotification(notif)
	go internal.SendEmailNotification(notif)
	return &pb.NotificationResponse{NotificationId: notif.ID.Hex()}, nil
}

func (s *NotificationServer) GetUserNotifications(ctx context.Context, req *pb.GetUserNotificationsRequest) (*pb.NotificationsList, error) {
	userId, _ := primitive.ObjectIDFromHex(req.UserId)
	filter := bson.M{"userId": userId}
	cursor, err := s.DB.Collection("notifications").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var notifs []models.Notification
	cursor.All(ctx, &notifs)
	var items []*pb.NotificationItem
	for _, n := range notifs {
		items = append(items, &pb.NotificationItem{
			Id:        n.ID.Hex(),
			Type:      string(n.Type),
			Message:   n.Message,
			TaskId:    n.TaskID.Hex(),
			CreatedAt: n.CreatedAt.Format(time.RFC3339),
			Read:      n.Read,
		})
	}
	return &pb.NotificationsList{Notifications: items}, nil
}
