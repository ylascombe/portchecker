package models

import "portchecker/gorm_custom"

type CheckAgentOutFlux struct {
	gorm_custom.GormModelCustom

	To string `json:"to" yaml:"to"`
	Port int `json:"port" yaml:"port"`

	Status bool `json:"status" yaml:"status"`
}

type TransformedCheckAgentOutFlux struct {
	ID uint       `gorm:"primary_key" yaml:"ID" json:"ID"`

	To string `json:"to" yaml:"to"`
	Port int `json:"port" yaml:"port"`
	Status bool `json:"status" yaml:"status"`
}

//func (inFlux CheckAgentInFlux) IsValid() bool {
//	return inFlux.UserID != 0 &&
//		inFlux.User.ID == inFlux.UserID &&
//		inFlux.EnvironmentID != 0 &&
//		inFlux.Environment.ID == inFlux.EnvironmentID
//}

func TransformCheckAgentOutFlux(inFlux []CheckAgentOutFlux) *[]TransformedCheckAgentOutFlux {
	var res []TransformedCheckAgentOutFlux

	for i:=0; i<len(inFlux); i++ {
		res = append(res, TransformedCheckAgentOutFlux{
			To: inFlux[i].To,
			Port: inFlux[i].Port,
			Status: inFlux[i].Status,
		})
	}

	return &res
}
