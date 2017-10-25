package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"portchecker/db_models"
)

var (
	email = "email@localhost"
)

func TestInsert(t *testing.T) {

	db := NewDBDriver()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&db_models.CheckAgent{})

	// Create
	db.Create(&db_models.CheckAgent{Hostname: "local"})

	// Read
	var agent db_models.CheckAgent
	res := db.First(&agent, "Hostname = ?", "local")

	assert.NotNil(t, res)

	// Delete
	//db.Delete(&agent)

}

func TestAutoMigrateDB(t *testing.T) {
	db := NewDBDriver()
	defer db.Close()

	AutoMigrateDB(db)
}

// Force to remove test user
func TestTearDown(t *testing.T) {
	// remove user in order to not change initial state
	db := NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	//res := db.Exec("delete from checkagent where email = ?", email).Error
	//assert.Nil(t, res)
}


func TestDataset(t *testing.T) {


	db := NewDBDriver()
	defer db.Close()

	// LOAD BALANCERS
	checkAgentInFluxLB1 := []db_models.CheckAgentInFlux{}

	checkAgentInFluxLB1 = append(checkAgentInFluxLB1, db_models.CheckAgentInFlux{
		Requested: true,
		Port: 80,
		From: "*",
	})

	checkAgentInFluxLB1 = append(checkAgentInFluxLB1, db_models.CheckAgentInFlux{
		Requested: true,
		Port: 443,
		From: "*",
	})

	checkAgentOutFluxLB1 := []db_models.CheckAgentOutFlux{}

	checkAgentOutFluxLB1 = append(checkAgentOutFluxLB1, db_models.CheckAgentOutFlux{
		Status: true,
		Port: 8080,
		To: "webserver1",
	})

	checkAgentOutFluxLB1 = append(checkAgentOutFluxLB1, db_models.CheckAgentOutFlux{
		Status: true,
		Port: 8080,
		To: "webserver2",
	})

	lb1 := db_models.CheckAgent{
		AnalysisId: 1,
		Hostname:   "lb1",
		InFlux:     checkAgentInFluxLB1,
		OutFlux:    checkAgentOutFluxLB1,
	}

	lb2 := db_models.CheckAgent{
		AnalysisId: 1,
		Hostname:   "lb2",
		InFlux:     checkAgentInFluxLB1,
		OutFlux:    checkAgentOutFluxLB1,
	}

	db.Create(&lb1)
	db.Create(&lb2)


	// WEBSERVERS
	checkAgentInFluxWebServer := []db_models.CheckAgentInFlux{}

	checkAgentInFluxWebServer = append(checkAgentInFluxWebServer, db_models.CheckAgentInFlux{
		Requested: true,
		Port: 8080,
		From: "lb1",
	})

	checkAgentInFluxWebServer = append(checkAgentInFluxWebServer, db_models.CheckAgentInFlux{
		Requested: true,
		Port: 8080,
		From: "lb2",
	})

	checkAgentOutFluxWebserver := []db_models.CheckAgentOutFlux{}

	checkAgentOutFluxWebserver = append(checkAgentOutFluxWebserver, db_models.CheckAgentOutFlux{
		Status: true,
		Port: 3306,
		To: "db-master",
	})
	webserver1 := db_models.CheckAgent{
		AnalysisId: 1,
		Hostname:   "webserver1",
		InFlux:     checkAgentInFluxWebServer,
		OutFlux:    checkAgentOutFluxWebserver,
	}

	webserver2 := db_models.CheckAgent{
		AnalysisId: 1,
		Hostname:   "webserver2",
		InFlux:     checkAgentInFluxWebServer,
		OutFlux:    checkAgentOutFluxWebserver,
	}

	// Migrate the schema
	db.Create(&webserver1)
	db.Create(&webserver2)


	// DATABASE
	checkAgentInFluxDb := []db_models.CheckAgentInFlux{}

	checkAgentInFluxDb = append(checkAgentInFluxDb, db_models.CheckAgentInFlux{
		Requested: true,
		Port: 3306,
		From: "webserver1",
	})

	checkAgentInFluxWebServer = append(checkAgentInFluxWebServer, db_models.CheckAgentInFlux{
		Requested: true,
		Port: 3306,
		From: "webserver2",
	})

	checkAgentOutFluxDb := []db_models.CheckAgentOutFlux{}

	checkAgentOutFluxDb = append(checkAgentOutFluxDb, db_models.CheckAgentOutFlux{
		Status: true,
		Port: 9090,
		To: "prometheus",
	})

	checkAgentDb := db_models.CheckAgent{
		AnalysisId: 1,
		Hostname:   "db-master",
		InFlux:     checkAgentInFluxDb,
		OutFlux:    checkAgentOutFluxDb,
	}

	db.Create(&checkAgentDb)


}
