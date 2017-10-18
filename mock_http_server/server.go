package mock_http_server

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"portchecker/models"
	"net"
)


func CreateHTTPListenServer(port int, timeout int, channel chan models.MockServerResult) {

	srv := &http.Server{Addr: fmt.Sprintf(":%v", port)}

	handleHello := sayHello(channel)
	http.HandleFunc("/", handleHello)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	time.Sleep(time.Duration(timeout) * time.Second)
	channel <- models.MockServerResult{Status:-1}
	srv.Shutdown(nil)
}

func sayHello(channel chan models.MockServerResult) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		channel <- models.MockServerResult{Status:0}
		fmt.Fprintf(w, "Hello")
	}
}

func CreateTCPListenServer(port int, timeout int, channel chan models.MockServerResult) {

	listenPort := fmt.Sprintf(":%v", port)
	fmt.Println("ListenPort:", listenPort)
	// Listen for incoming connections.
	listener, err := net.Listen("tcp", listenPort)

	// Close the listener when the application closes.
	// defer listener.Close()
	fmt.Println("Listening on " + listenPort)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		channel <- models.MockServerResult{Status:-1}
		return
	}

	go func() {
		for {
			// Listen for an incoming connection.
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				channel <- models.MockServerResult{Status:-1}
				return
			}
			// Handle connections in a new goroutine.
			go handleTCPRequest(conn, channel)
		}
	}()

	time.Sleep(time.Duration(timeout) * time.Second)
	listener.Close()
	channel <- models.MockServerResult{Status:-1}

}


func handleTCPRequest(conn net.Conn, channel chan models.MockServerResult) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()

	channel <- models.MockServerResult{Status:0}
}
