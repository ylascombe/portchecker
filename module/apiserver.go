package module

import (
	"portchecker/database"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"portchecker/controllers"
	"fmt"
	"portchecker/conf"
)

func StartApiServer(config conf.Config) {
	db := database.NewDBDriver()
	database.AutoMigrateDB(db)

	router := gin.New()
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))


	hostname := router.Group("/v1/check_agents")
	{
		hostname.GET("/", controllers.FetchAllCheckAgent)
	}

	probe_agents := router.Group("/v1/probe_agents")
	{
		probe_agents.GET("/", controllers.FetchAllProbeAgent)
	}

	createAgentReport := router.Group("/v1/hostname/:hostname/analysis_id/:analysis_id/check_agents")
	{
		createAgentReport.POST("/", controllers.CreateCheckAgentReport)
	}

	createProbeAgentReport := router.Group("/v1/hostname/:hostname/analysis_id/:analysis_id/probe_agent")
	{
		createProbeAgentReport.POST("/", controllers.CreateProbeAgentReport)
	}

	report := router.Group("/v1/report/:analysis_id")
	{
		report.GET("/", controllers.ExtractReport)
	}

	router.Run(fmt.Sprintf(":%v", config.ApiServerListenPort))

}
