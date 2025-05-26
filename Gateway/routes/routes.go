package routes

import (
	"gateway/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.LoadHTMLGlob("frontend/templates/*")
	// router.Use()
	// Attach auth middleware here
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
		analytics := router.Group("/analytics")
		analytics.GET("/user/:userID", controller.GetUserTaskStats)
		analytics.GET("/team/:teamID", controller.GetTeamTaskStats)
	}

	{
		teams := router.Group("/teams")
		teams.DELETE("/:id", controller.DeleteTeam)
		teams.POST("", controller.CreateTeam)
		teams.GET("/:id", controller.GetTeam)
		teams.GET("", controller.GetTeams)
		teams.PATCH("/:id", controller.UpdateTeam)
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
