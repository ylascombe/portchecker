package models

type VisjsGraph struct {
	Nodes []VisjsNode `json:"nodes" yaml:"nodes"`
	Edges []VisjsEdge `json:"edges" yaml:"edges"`

}
