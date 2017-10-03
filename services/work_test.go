package services

import (
	"testing"
	"portchecker/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
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
