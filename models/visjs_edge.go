package models

type VisjsEdge struct {
	From int `json:"from" yaml:"from"`
	To int `json:"to" yaml:"to"`
	Color string `json:"color" yaml:"color"`
	Label string `json:"label" yaml:"label"`
	Arrows string `json:"arrows" yaml:"arrows"`
}
