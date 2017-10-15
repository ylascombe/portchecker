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
