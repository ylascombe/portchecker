package models

type VisjsEdge struct {
	From string `json:"from" yaml:"from"`
	To string `json:"to" yaml:"to"`
	Color string `json:"color" yaml:"color"`
	Label string `json:"label" yaml:"label"`
}
