package routes

import (
	"gateway/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.LoadHTMLGlob("frontend/templates/*")
	router.Static("/static", "frontend/static")
	{
		router.GET("", controller.GetIndex)
	}

	{
		// users := router.Group("/users")
	}

	{
		// tasks := router.Group("/tasks")
	}

	{
		// analytics := router.Group("/analytics")
	}

	{
		// teams := router.Group("/teams")
	}

	{
		// notifications := router.Group("/notifications")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router.Run(":" + port)	
}