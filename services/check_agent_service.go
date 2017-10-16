package services

import (
	"portchecker/database"
	"portchecker/db_models"
	"portchecker/models"
	"fmt"
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

func FetchAllReportForAnalysis(analysisId int) (*[]models.ReportResult, error) {

	db := database.NewDBDriver()
	defer db.Close()

	var checkAgents []db_models.CheckAgent
	err := db.Where("analysis_id = ?", analysisId).Find(&checkAgents).Error

	if err != nil {
		return nil, err
	}

	var reportResults = []models.ReportResult{}
	for _, item := range checkAgents {

		status := true
		statusDetail := ""

		var outFluxes = []db_models.CheckAgentOutFlux{}
		db.Model(&item).Related(&outFluxes)

		for _, outFlux := range outFluxes {

			status = status && outFlux.Status
			if ! outFlux.Status {
				statusDetail = fmt.Sprintf("%v, Out flux KO %v", statusDetail, outFlux)
			}
		}

		var inFluxes = []db_models.CheckAgentInFlux{}
		db.Model(&item).Related(&inFluxes)

		for _, inFlux := range inFluxes {
			status = status && inFlux.Requested
			if ! inFlux.Requested {
				statusDetail = fmt.Sprintf("%v, In flux KO %v", statusDetail, inFlux)
			}
		}
		reportResults = append(reportResults, models.ReportResult{
			Hostname: item.Hostname,
			Status: status,
			StatusDetails: statusDetail,
		})
	}

	return &reportResults, nil

}
