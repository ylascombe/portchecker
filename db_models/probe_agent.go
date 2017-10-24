package db_models

import (
	"portchecker/gorm_custom"
)

type ProbeAgent struct {
	gorm_custom.GormModelCustom

	Hostname  string `json:"hostname" yaml:"hostname"`
	AnalysisId int `json:"analysis_id" yaml:"analysis_id"`

	OpenedPorts []int `gorm:"-" json:"opened_ports" yaml:"opended_ports"`
	OpenedPortsString string `json:"opened_ports_string" yaml:"opened_ports_string"`
}

type TransformedProbeAgent struct {
	ID   uint `json:"id"`

	Hostname  string `json:"hostname" yaml:"hostname"`
	AnalysisId int `json:"analysis_id" yaml:"analysis_id"`
	OpenedPorts string `json:"opened_ports" yaml:"opended_ports"`
}


func TransformprobeAgent(agent ProbeAgent) *TransformedProbeAgent {
	return &TransformedProbeAgent{
		ID: agent.ID,
		Hostname: agent.Hostname,
		OpenedPorts: agent.OpenedPortsString,
	}
}
