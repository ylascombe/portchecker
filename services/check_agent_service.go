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


func FetchAllReportForAnalysisV2(analysisId int) (*models.VisjsGraph, error) {

	db := database.NewDBDriver()
	defer db.Close()

	var checkAgents []db_models.CheckAgent
	err := db.Where("analysis_id = ?", analysisId).Find(&checkAgents).Error

	if err != nil {
		return nil, err
	}

	nodes := make(map[string]models.VisjsNode)
	edges := make(map[string]models.VisjsEdge)

	for _, item := range checkAgents {


		nodes = addNodeIfNotExist(nodes, item.Hostname)

		var outFluxes = []db_models.CheckAgentOutFlux{}
		db.Model(&item).Related(&outFluxes)

		for _, outFlux := range outFluxes {

			edgeKey := fmt.Sprintf("%v-%v-%v", item.Hostname, outFlux.To, outFlux.Port)
			_, ok := edges[edgeKey]

			color := "#31B404" // green
			if ! outFlux.Status {
				color = "#DF0101" // red
			}
			if !ok {
				edges[edgeKey] = models.VisjsEdge{
					To: nodes[outFlux.To].Id,
					From: nodes[item.Hostname].Id,
					Label: fmt.Sprintf("%v", outFlux.Port),
					Color: models.VisjsColor{
						Color: color,
						Inherit: "from",
						Opacity: "1.0",
					},
					Arrows: "to",
				}
			}

			nodes = addNodeIfNotExist(nodes, outFlux.To)
		}

		var inFluxes = []db_models.CheckAgentInFlux{}
		db.Model(&item).Related(&inFluxes)

		for _, inFlux := range inFluxes {
			edgeKey := fmt.Sprintf("%v-%v-%v", inFlux.From, item.Hostname, inFlux.Port)
			_, ok := edges[edgeKey]

			color := "#31B404" // green
			if ! inFlux.Requested {
				color = "#DF0101" // red
			}
			if !ok {
				edges[edgeKey] = models.VisjsEdge{
					To: nodes[item.Hostname].Id,
					From: nodes[inFlux.From].Id,
					Label: fmt.Sprintf("%v", inFlux.Port),
					Color: models.VisjsColor{
						Color: color,
						Inherit: "from",
						Opacity: "1.0",
					},
					Arrows: "to",
				}
			}

			nodes = addNodeIfNotExist(nodes, inFlux.From)
		}

	}

	arrayNode := []models.VisjsNode{}
	for _, value := range nodes {
		arrayNode = append(arrayNode, value)
	}

	arrayEdges := []models.VisjsEdge{}
	for _, value := range edges {
		arrayEdges = append(arrayEdges, value)
	}

	graph := models.VisjsGraph{
		Nodes: arrayNode,
		Edges: arrayEdges,
	}

	return &graph, nil
}

func addNodeIfNotExist(nodes map[string]models.VisjsNode, nodeName string) map[string]models.VisjsNode {
	_, ok := nodes[nodeName]

	// add node in map if still does not exists
	if ! ok {
		nodes[nodeName] = models.VisjsNode{
			Id: len(nodes),
			Label: nodeName,
			Title: nodeName,
		}
	}

	return nodes
}
