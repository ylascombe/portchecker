package models

type CheckResult struct {
	ActionList ActionList
	NotRequestedPort []int
	NotFunctionnalOutFlux []Route
}

func (checkResult CheckResult) ToString() string {
	return "TODO"
}
