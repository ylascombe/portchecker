package services

import (
	"portchecker/models"
	"portchecker/database"
)

func ListCheckAgent() (*[]models.CheckAgent, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var checkAgent []models.CheckAgent
	err := db.Find(&checkAgent).Error

	return &checkAgent, err
}
