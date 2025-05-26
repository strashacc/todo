package controller

import (
	"gateway/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserTaskStats(c *gin.Context) {
	userID := c.Param("userID")
	res, err := repo.GetUserTaskStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_tasks":       res.TotalTasks,
		"completed_tasks":   res.CompletedTasks,
		"in_progress_tasks": res.InProgressTasks,
		"overdue_tasks":     res.OverdueTasks,
	})
}

func GetTeamTaskStats(c *gin.Context) {
	teamID := c.Param("teamID")
	res, err := repo.GetTeamTaskStats(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_tasks":       res.TotalTasks,
		"completed_tasks":   res.CompletedTasks,
		"in_progress_tasks": res.InProgressTasks,
		"overdue_tasks":     res.OverdueTasks,
	})
}
