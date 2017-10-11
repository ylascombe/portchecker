package services

import (
	"testing"
	"portchecker/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"portchecker/models"
	"time"
	"net"
)

func TestMakeActionListHostnameNotConcerned(t *testing.T) {

	// arrange
	content, _ := utils.UnmarshallFromFile("../mapping.yml")
	fmt.Println(content)

	// act
	res, err := MakeActionList(*content, "not-concerned")

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 0, len(res.ListenOnPort))
	assert.Equal(t, 0, len(res.TestFlux))

}

func TestMakeActionList(t *testing.T) {

	// arrange
	content, _ := utils.UnmarshallFromFile("../mapping.yml")
	fmt.Println(content)

	// act
	res, err := MakeActionList(*content, "vm1-vlan1")

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(res.ListenOnPort))
	assert.Equal(t, 22, res.ListenOnPort[0])

	assert.Equal(t, 2, len(res.TestFlux))

	assert.Equal(t, "vm1-vlan1", res.TestFlux[0].From)
	assert.Equal(t, "vm1-vlan2", res.TestFlux[0].To)
	assert.Equal(t, 9200, res.TestFlux[0].Port)

	assert.Equal(t, "vm1-vlan1", res.TestFlux[1].From)
	assert.Equal(t, "vm1-vlan2", res.TestFlux[1].To)
	assert.Equal(t, 9900, res.TestFlux[1].Port)

}

func TestCreateMockServers(t *testing.T) {
	// arrange
	actionList := models.ActionList{ListenOnPort:[]int{22, 80, 90}}
	checkResult := models.CheckResult{
		ActionList: actionList,
	}
	timeout := 5

	mockFunc := func (port int, timeout int, channel chan models.MockServerResult)  {
		// fmt.Println("wait ", port, "seconds")

		time.Sleep(time.Duration(port) * time.Millisecond )

		if port == 90 {
		channel <- models.MockServerResult{Port: port, Status: -1}
		}
		// fmt.Println("fin ", port)

		channel <- models.MockServerResult{Port: port, Status: 0}
	}
	// act
	err := CreateMockServers(&checkResult, timeout, mockFunc)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(checkResult.NotRequestedPort))
	assert.Equal(t, 90, checkResult.NotRequestedPort[0])
}


func TestTestFlux(t *testing.T) {
	// arrange
	res, _ := net.LookupHost("google.fr")

	fmt.Println(res)
	routesToTest := []models.Route{}

	existingRoute := models.Route{To: res[1], Port: 80}
	routesToTest = append(routesToTest, existingRoute)

	unexistingRoute  := models.Route{To: "127.0.0.1", Port: 1}
	routesToTest = append(routesToTest, unexistingRoute)

	actionList := models.ActionList{TestFlux:routesToTest}
	checkResult := models.CheckResult{
		ActionList: actionList,
	}

	// act
	err := TestFlux(&checkResult)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(checkResult.NotFunctionnalOutFlux))
}
