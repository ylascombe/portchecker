package models

import "fmt"

type Route struct {

	From string `json:"from" yaml:"from"`
	To string `json:"to" yaml:"to"`
	Port int `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

func (route *Route) ToString() string {
	return fmt.Sprintf("from: %v, to: %v, port: %v, mode: %v", route.From, route.To, route.Port, route.Mode)
}
