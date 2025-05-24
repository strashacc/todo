package controller

import (
	"gateway/pb/team_pb"
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

func CreateTeam(c *gin.Context) {
	var team team_pb.CreateTeamRequest
	c.Bind(&team)
	res, err := repo.CreateTeam(&team)
	if err != nil {
		log.Println(err.Error())
	}
	if !res.Success {
		c.JSON(200, res)
		return
	}
	c.JSON(201, res)
}
