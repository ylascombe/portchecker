package models

import "fmt"

type CheckResult struct {
	ActionList ActionList
	NotRequestedPort []int
	NotFunctionnalOutFlux []Route
}

func (checkResult CheckResult) ToString() string {
	return fmt.Sprintf("CheckResult{ActionList: , NotRequestedPort, %v, NotFunctionnalOutFlux: %v}",
		checkResult.ActionList,
		checkResult.NotRequestedPort,
		checkResult.NotFunctionnalOutFlux)
}
func (checkResult CheckResult) PrintResult() string {
	return fmt.Sprintf("CheckResult{NotRequestedPort, %v, NotFunctionnalOutFlux: %v}",
		checkResult.NotRequestedPort,
		checkResult.NotFunctionnalOutFlux)
}

func (checkResult CheckResult) ErrorNumber() int {
	return len(checkResult.NotFunctionnalOutFlux) + len(checkResult.NotRequestedPort)
}
