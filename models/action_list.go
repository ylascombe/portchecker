package models

import "portchecker/gorm_custom"

type ActionList struct {
	gorm_custom.GormModelCustom

	ListenOnPort []int
	TestFlux []Route
}
