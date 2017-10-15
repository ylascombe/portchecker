package db_models

import "portchecker/gorm_custom"

type CheckAgent struct {
	gorm_custom.GormModelCustom

	Hostname  string `json:"hostname" yaml:"hostname"`
	InFlux []CheckAgentInFlux `json:"in_flux" yaml:"in_flux"`
	OutFlux []CheckAgentOutFlux `json:"out_flux" yaml:"out_flux"`
}

type TransformedCheckAgent struct {
	ID   uint `json:"id"`

	Hostname  string `json:"hostname" yaml:"hostname"`
	InFlux []TransformedCheckAgentInFlux `json:"in_flux" yaml:"in_flux"`
	OutFlux []TransformedCheckAgentOutFlux `json:"out_flux" yaml:"out_flux"`
}

func (checkAgent CheckAgent) IsValid() bool {
	return checkAgent.Hostname != ""
}

func TransformCheckAgent(agent CheckAgent) *TransformedCheckAgent {
	return &TransformedCheckAgent{
		ID: agent.ID,
		Hostname: agent.Hostname,
		InFlux: *TransformCheckAgentInFlux(agent.InFlux),
		OutFlux: *TransformCheckAgentOutFlux(agent.OutFlux),
	}
}
