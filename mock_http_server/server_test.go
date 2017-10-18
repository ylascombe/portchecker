package mock_http_server

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"portchecker/models"
	"fmt"
	"net/http"
	"net"
	"bufio"
)

func TestCreateHTTPListenServerWhenNoConnection(t *testing.T) {
	// arrange
	channel := make(chan models.MockServerResult,1)
	timeout := 2

	// act
	CreateHTTPListenServer(9090, timeout, channel)

	// assert
	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, -1, res.Status)
}

func TestCreateHTTPListenServerWhenConnection(t *testing.T) {
	// arrange
	channel := make(chan models.MockServerResult,1)
	timeout := 5
	port := 9091
	url := fmt.Sprintf("http://localhost:%v/", port)

	// act
	go CreateHTTPListenServer(port, timeout, channel)

	go http.Get(url)

	// assert
	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, 0, res.Status)
}

func TestCreateTCPListenServerWhenNoConnection(t *testing.T) {
	// arrange
	channel := make(chan models.MockServerResult,1)
	timeout := 2

	// act
	CreateTCPListenServer(9090, timeout, channel)

	// assert
	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, -1, res.Status)
}

func TestCreateTCPListenServerWhenOneConnection(t *testing.T) {

	// arrange
	channel := make(chan models.MockServerResult,1)
	timeout := 2

	// act
	go CreateTCPListenServer(9123, timeout, channel)

	conn, _ := net.Dial("tcp", ":9123")
	// send to socket
	fmt.Fprintf(conn, "Who are you ?")
	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

	// assert
	res := <- channel
	assert.NotNil(t, res)
	assert.Equal(t, 0, res.Status)
	assert.Equal(t, "Message received.", message)
}
