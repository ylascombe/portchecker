package services

import (
	"portchecker/models"
	"fmt"
	"errors"
	"log"
	"net/http"
)

func DoWork(config models.Config, hostname string) (int, error) {


	actionList, err := MakeActionList(config, hostname)

	// TODO variabilize timeout delay
	timeout := 120

	if err != nil {
		return -1, err
	}

	checkResult := models.CheckResult{ActionList: actionList}
	err = TestFlux(actionList, &checkResult)

	if err != nil {
		return -1, err
	}
	err = CreateMockServers(&checkResult, timeout, CreateListenServer)

	if err != nil {
		return -1, err
	}
	nbErrors := len(checkResult.NotFunctionnalOutFlux) + len(checkResult.NotRequestedPort)
	if nbErrors > 0 {
		return nbErrors, errors.New(fmt.Sprintf("At least an error has occured, errors details : %v", checkResult.ToString()))
	}
	return 0, nil
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

func TestFlux(actionList models.ActionList, checkResult *models.CheckResult) error {
	// TODO
	return nil
}


func CreateListenServer(port int, timeout int, channel chan models.MockServerResult) {
	http.HandleFunc("/", sayHello) // set router
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
