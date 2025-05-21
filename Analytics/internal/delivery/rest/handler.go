package rest

import (
	"net/http"

	"analytics/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRESTHandler(router *gin.Engine, uc usecase.TaskUsecase) {
	router.GET("/api/stats", func(c *gin.Context) {
		stats, err := uc.GetStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stats)
	})
}
