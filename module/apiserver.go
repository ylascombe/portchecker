package module

import (
	"portchecker/database"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"portchecker/controllers"
)

func StartApiServer() {
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


	hostname := router.Group("/v1/hostname/:hostname/check_agents")
	{
		hostname.GET("/", controllers.FetchAllCheckAgent)
		hostname.POST("/", controllers.CreateCheckAgentReport)
	}

	router.Run(":8090")

}
