package services

import (
	"portchecker/database"
	"portchecker/db_models"
)

func ListCheckAgent() (*[]db_models.CheckAgent, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var checkAgent []db_models.CheckAgent
	err := db.Find(&checkAgent).Error

	return &checkAgent, err
}

func CreateCheckAgentReport(toAdd db_models.CheckAgent) (*db_models.CheckAgent, error) {
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
