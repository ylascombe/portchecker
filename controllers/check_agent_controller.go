package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portchecker/db_models"
	"portchecker/services"
)

func FetchAllCheckAgent(c *gin.Context) {
	var _checkAgents []db_models.TransformedCheckAgent

	checkAgents, err := services.ListCheckAgent()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status" : http.StatusInternalServerError,
			"error message" : err,
		})
		return
	}

	if (len(*checkAgents) <= 0) {
		// choice : if no item found, return a HTTP status code 200 with an empty array
		_checkAgents = make([]db_models.TransformedCheckAgent, 0)
	}

	//transforms check agent
	for _, item := range *checkAgents {
		tmp := db_models.TransformCheckAgent(item)
		_checkAgents = append(_checkAgents, *tmp)
	}
	c.JSON(http.StatusOK, _checkAgents)
}
