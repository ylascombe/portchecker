package module

import (
	"portchecker/models"
	"fmt"
	"errors"
	"portchecker/mock_http_server"
	"net/http"
	"io/ioutil"
	"portchecker/db_models"
)

func StartCheckAgent(config models.Config, hostname string, timeout int) (*models.CheckResult, error) {

	actionList, err := MakeActionList(config, hostname)
	if err != nil {
		return nil, err
	}

	checkResult := models.CheckResult{ActionList: actionList}

	err = CreateMockServers(&checkResult, timeout, mock_http_server.CreateListenServer)
	if err != nil {
		return nil, err
	}

	err = TestFlux(&checkResult)
	if err != nil {
		return nil, err
	}

	if checkResult.ErrorNumber() > 0 {
		return &checkResult, errors.New(
			fmt.Sprintf("%v error has occured, errors details : %v",
			checkResult.ErrorNumber(),
			checkResult.PrintResult(),
		))
	}
	return &checkResult, nil
}

func MakeActionList(config models.Config, hostname string) (models.ActionList, error) {

	fmt.Println("hostname: ", hostname)

	var actionList = models.ActionList{}

	for i:=0; i<len(config.Routes); i++ {

		// simulate mock server for port where listening in mode : check
		if config.Routes[i].To == hostname {
			fmt.Println("Create mock server for listening on ", config.Routes[i].ToString())
			actionList.ListenOnPort = append(actionList.ListenOnPort, config.Routes[i].Port)
		}

		// test flux for route starting from current host
		if config.Routes[i].From == hostname {
			fmt.Println("Test connection ", config.Routes[i].ToString())
			actionList.TestFlux = append(actionList.TestFlux, config.Routes[i])
		}

	}
	return actionList, nil
}

func CreateMockServers(checkResult *models.CheckResult, timeout int, mockFunc func(port int, timeout int, channel chan models.MockServerResult)) error {

	var channels = make(chan models.MockServerResult, len(checkResult.ActionList.ListenOnPort))

	for i:=0; i<len(checkResult.ActionList.ListenOnPort); i++ {
		go mockFunc(checkResult.ActionList.ListenOnPort[i], timeout, channels)
	}

	for i:=0; i<len(checkResult.ActionList.ListenOnPort); i++ {
		mockServerRes := <- channels

		if mockServerRes.Status != 0 {
			checkResult.NotRequestedPort = append(checkResult.NotRequestedPort, mockServerRes.Port)
		}
	}
	return nil
}

func TestFlux(checkResult *models.CheckResult) error {

	for i:=0; i<len(checkResult.ActionList.TestFlux); i++ {
		route := checkResult.ActionList.TestFlux[i]

		url := fmt.Sprintf("http://%v:%v", route.To, route.Port)
		resp, err := http.Get(url)

		if err != nil {
			checkResult.NotFunctionnalOutFlux = append(checkResult.NotFunctionnalOutFlux, route)
		} else {
			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				checkResult.NotFunctionnalOutFlux = append(checkResult.NotFunctionnalOutFlux, route)
			}
		}
	}

	return nil
}

// TODO add test
func ProcessCheckAgentResult(config models.Config, checkResult models.CheckResult, hostname string) (*db_models.CheckAgent, error) {

	checkAgent := db_models.CheckAgent{}

	for _, route := range config.Routes {

		if route.To == hostname {
			// IN FLUX
			checkAgent.InFlux = append(checkAgent.InFlux,
				db_models.CheckAgentInFlux{
					From:      route.From,
					Port:      route.Port,
					Requested: portInSlice(route.Port, checkResult.NotRequestedPort),
				},
			)
		} else {
			// OUT FLUX
			found := false
			for i:=0; i<len(checkResult.NotFunctionnalOutFlux) && !found; i++ {
				found = checkResult.NotFunctionnalOutFlux[i].To == hostname
			}

			checkAgent.OutFlux = append(checkAgent.OutFlux,
				db_models.CheckAgentOutFlux{
					To: route.To,
					Port: route.Port,
					Status: found,
				},
			)
		}

	}
	return &checkAgent, nil

}

func portInSlice(port int, list []int) bool {
	for _, item := range list {
		if item == port {
			return true
		}
	}
	return false
}
