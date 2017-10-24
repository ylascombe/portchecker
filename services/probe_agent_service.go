package services

import (
	"portchecker/database"
	"portchecker/db_models"
)

func ListProbeAgent() (*[]db_models.ProbeAgent, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var probeAgents []db_models.ProbeAgent
	err := db.Find(&probeAgents).Error

	return &probeAgents, err
}

func CreateProbeAgentReport(toAdd db_models.ProbeAgent) (*db_models.ProbeAgent, error) {
	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&toAdd).Error

	if err == nil {
		err = db.First(&toAdd).Error
	}

	if err != nil {
		return nil, err
	} else {
		return &toAdd, nil
	}
}
