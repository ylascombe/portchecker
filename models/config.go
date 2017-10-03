package models

type Config struct {
	FormatVersion string           `json:"format_version" yaml:"format_version"`

	Routes []Route `json:"routes" yaml:"routes"`
	MainParams MainParams  `json:"main_params" yaml:"main_params"`
}
