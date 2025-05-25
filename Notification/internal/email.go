package internal

import (
    "fmt"
    "github.com/strashacc/todo/Notification/models"
)

func SendEmailNotification(notif models.Notification) {
    fmt.Printf("[EMAIL] Notification: %s (UserID: %v, Message: %s)\n", notif.ID.Hex(), notif.UserID, notif.Message)
}
