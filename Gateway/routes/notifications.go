package routes

import (
	"gateway/controller"
	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(r *gin.Engine) {
	r.GET("/api/notifications", controller.GetNotifications)
}
