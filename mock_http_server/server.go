package mock_http_server

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"portchecker/models"
)

// TODO Improve this class, currently, this is a experimental one

//func handler(w http.ResponseWriter, r *http.Request) {
//	hostname, _ := os.Hostname()
//	fmt.Fprintf(w, "Hi, I'm %s", hostname)
//}
//
//func niceHello(w http.ResponseWriter, r *http.Request) {
//	hostname, _ := os.Hostname()
//	fmt.Fprintf(w, "Hi %s, I'm %s", r.URL.Path[1:], hostname)
//}
//
//func main() {
//	http.HandleFunc("/", niceHello)
//	http.ListenAndServe(":8080", nil)
//}


func CreateListenServer(port int, timeout int, channel chan models.MockServerResult) {

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
