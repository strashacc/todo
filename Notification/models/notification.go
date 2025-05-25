package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationType string

const (
	NotificationGlobal   NotificationType = "global"
	NotificationTeam     NotificationType = "team"
	NotificationPersonal NotificationType = "personal"
)

type Notification struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Type      NotificationType    `bson:"type"`
	UserID    *primitive.ObjectID `bson:"userId,omitempty"`
	TeamID    *primitive.ObjectID `bson:"teamId,omitempty"`
	Message   string              `bson:"message"`
	TaskID    primitive.ObjectID  `bson:"taskId"`
	Read      bool                `bson:"read"`
	CreatedAt time.Time           `bson:"createdAt"`
}
