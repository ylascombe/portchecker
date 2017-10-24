package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portchecker/db_models"
	"portchecker/services"
	"fmt"
	"strconv"
)

func FetchAllProbeAgent(c *gin.Context) {
	var _probeAgents []db_models.TransformedProbeAgent

	probeAgents, err := services.ListProbeAgent()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status" : http.StatusInternalServerError,
			"error message" : err,
		})
		return
	}

	if (len(*probeAgents) <= 0) {
		// choice : if no item found, return a HTTP status code 200 with an empty array
		_probeAgents = make([]db_models.TransformedProbeAgent, 0)
	}

	//transforms probe agent
	for _, item := range *probeAgents {
		tmp := db_models.TransformprobeAgent(item)
		_probeAgents = append(_probeAgents, *tmp)
	}
	c.JSON(http.StatusOK, _probeAgents)
}


func CreateProbeAgentReport(c *gin.Context) {

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
	fmt.Println("Create probeAgent result for hostname", hostname)
	var probeAgentJSON db_models.ProbeAgent

	err = c.BindJSON(&probeAgentJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"message" : "Invalid request.",
			"error detail": err,
		})
	} else {
		probeAgentJSON.Hostname = hostname

		probeAgentJSON.AnalysisId = analysisId
		probeAgent, err := services.CreateProbeAgentReport(probeAgentJSON)
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
				"Location":     fmt.Sprintf("/v1/hostname/%v/check_agent/%v", hostname, probeAgent.ID),
				"probeAgentID": probeAgent.ID,
			})
		}
	}
}
