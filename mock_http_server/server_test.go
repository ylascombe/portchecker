package mock_http_server

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
	timeout := 5
	port := 9091
	url := fmt.Sprintf("http://localhost:%v/", port)

	// act
	go CreateListenServer(port, timeout, channel)

	go http.Get(url)

	// assert
	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, 0, res.Status)
}
