package models

type VisjsColor struct {
	Color string `json:"color" yaml:"color"`
	Highlight string `json:"highlight" yaml:"highlight"`
	Hover string `json:"hover" yaml:"hover"`
	Inherit string `json:"inherit" yaml:"inherit"`
	Opacity string `json:"opacity" yaml:"opacity"`
}
