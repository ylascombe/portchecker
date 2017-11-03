package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portchecker/db_models"
	"portchecker/services"
	"fmt"
	"strconv"
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


func CreateCheckAgentReport(c *gin.Context) {

	hostname := c.Param("hostname")
	analysisIdStr := c.Param("analysis_id")
	analysisId, err := strconv.Atoi(analysisIdStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status" : http.StatusInternalServerError,
				"message" : "Invalid analysis_id parameter",
			},
		)
		return
	}
	fmt.Println("Create checkAgent result for hostname", hostname)
	var checkAgentJSON db_models.CheckAgent

	err = c.BindJSON(&checkAgentJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"message" : "Invalid request.",
			"error detail": err,
		})
	} else {
		checkAgentJSON.Hostname = hostname

		checkAgentJSON.AnalysisId = analysisId
		checkAgent, err := services.CreateCheckAgentReport(checkAgentJSON)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status" : http.StatusInternalServerError,
				"message" : "Error while creating check agent report", "error detail": err,
				},
			)
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"status" :      http.StatusCreated,
				"message" :     "Check agent report created successfully!",
				"Location":     fmt.Sprintf("/v1/hostname/%v/check_agent/%v", hostname, checkAgent.ID),
				"checkAgentID": checkAgent.ID,
			})
		}
	}
}

func ExtractReport(c *gin.Context) {
	analysisIdStr := c.Param("analysis_id")
	analysisId, err := strconv.Atoi(analysisIdStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status" : http.StatusInternalServerError,
				"message" : "Invalid analysis_id parameter",
			},
		)
		return
	}

	graph, err := services.FetchAllReportForAnalysisV2(analysisId)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status" : http.StatusInternalServerError,
				"message" : err,
			},
		)
		return
	}

	c.JSON(http.StatusOK, graph)
}
