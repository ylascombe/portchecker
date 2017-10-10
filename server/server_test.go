package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"portchecker/models"
	"fmt"
	"net/http"
)

func TestCreateListenServerWhenNoConnection(t *testing.T) {
	// arrange
	channel := make(chan models.MockServerResult,1)
	timeout := 2

	// act
	CreateListenServer(9090, timeout, channel)

	// assert
	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, -1, res.Status)
}


func TestCreateListenServerWhenConnection(t *testing.T) {
	// arrange
	channel := make(chan models.MockServerResult,1)
	timeout := 20
	port := 9091
	url := fmt.Sprintf("http://localhost:%v/", port)
	fmt.Println(url)
	// act
	CreateListenServer(port, timeout, channel)

	resp, err := http.Get(url)

	// assert
	assert.Nil(t, err)

	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, 0, res.Status)
	assert.Equal(t, "", resp)
}
