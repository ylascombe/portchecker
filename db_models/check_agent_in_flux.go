package db_models

import "portchecker/gorm_custom"

type CheckAgentInFlux struct {
	gorm_custom.GormModelCustom

	CheckAgentID uint
	From string `json:"from" yaml:"from"`
	Port int `json:"port" yaml:"port"`
	Requested bool `json:"requested" yaml:"requested"`
}

type TransformedCheckAgentInFlux struct {
	ID        uint       `gorm:"primary_key" yaml:"ID" json:"ID"`

	From string `json:"from" yaml:"from"`
	Port int `json:"port" yaml:"port"`
	Requested bool `json:"requested" yaml:"requested"`
}

//func (inFlux CheckAgentInFlux) IsValid() bool {
//	return inFlux.UserID != 0 &&
//		inFlux.User.ID == inFlux.UserID &&
//		inFlux.EnvironmentID != 0 &&
//		inFlux.Environment.ID == inFlux.EnvironmentID
//}

func TransformCheckAgentInFlux(inFlux []CheckAgentInFlux) *[]TransformedCheckAgentInFlux {
	var res []TransformedCheckAgentInFlux

	for i:=0; i<len(inFlux); i++ {
		res = append(res,
			TransformedCheckAgentInFlux{
				From: inFlux[i].From,
				Port: inFlux[i].Port,
				Requested: inFlux[i].Requested,
			},
		)
	}

	return &res
}
