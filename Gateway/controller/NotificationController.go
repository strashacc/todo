package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Notification структура для ответа
type Notification struct {
	Text string `json:"text"`
}

// Handler для получения уведомлений
func GetNotifications(c *gin.Context) {
	notifications := []Notification{
		{Text: "Добро пожаловать!"},
		{Text: "У тебя новое задание."},
		{Text: "Пора делать перерыв!"},
	}
	c.JSON(http.StatusOK, notifications)
}
