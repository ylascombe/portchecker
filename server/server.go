package main

import (
	"fmt"
	"net/http"
	"os"
)

// TODO Improve this class, currently, this is a experimental one

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hi, I'm %s", hostname)
}

func niceHello(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hi %s, I'm %s", r.URL.Path[1:], hostname)
}

func main() {
	http.HandleFunc("/", niceHello)
	http.ListenAndServe(":8080", nil)
}
