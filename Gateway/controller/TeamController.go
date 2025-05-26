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

func GetTeam(c *gin.Context) {
	teamId := c.Param("id")
	// add authorization	

	res, err := repo.GetTeam(teamId)
	if res.Id == "" {
		c.JSON(203, team_pb.TeamResponse{Success: false, Message: "Team not found"})
		return
	} else if err != nil {
		c.JSON(203, res)
		return
	}
	c.JSON(200, res)
}

func GetTeams(c *gin.Context) {
	res, err := repo.GetTeams()
	if err != nil {
		c.JSON(203, team_pb.TeamResponse{Success: false, Message: "Couldn't get team list"})
		return
	} else if len(res.Teams) == 0 {
		c.JSON(200, team_pb.TeamResponse{Success: true, Message: "No teams found"})
		return
	}
	c.JSON(200, res)
}

func UpdateTeam(c *gin.Context) {
	teamId := c.Param("id")	
	var team team_pb.UpdateTeamRequest
	team.TeamId = teamId
	c.Bind(team.Update)
	// add authorization with c.get('userId') and verification of admin rights
	res, err := repo.UpdateTeam(&team)
	if err != nil {
		log.Println("Failed team update: " + err.Error())
		c.JSON(203, team_pb.TeamResponse{Success: false, Message: "Error"})
		return
	} else if res.Success == false {
		c.JSON(203, res)
		return
	}
	c.JSON(201, res)
}