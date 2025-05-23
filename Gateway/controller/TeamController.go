package controller

import (
	"gateway/repo"
	"log"

	"github.com/gin-gonic/gin"
)

func DeleteTeam(c *gin.Context) {
	token := c.Request.Header["token"]
	teamId := c.Param("id")
	token = token
	// Add authorization
	res, err := repo.DeleteTeam(teamId)
	log.Println(res)
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, res)
}
