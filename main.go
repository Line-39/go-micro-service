package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there 👋"))
}

func main() {
	// create a new servermux
	mux := http.NewServeMux()
	// register hello() as a handler for "/" pattern
	mux.HandleFunc("/", hello)

	// log the service startup
	log.Print("starting service on :4000")

	// start http server on :4000
	err := http.ListenAndServe(":4000", mux)
	// log the error message if ListenAndServe() encounters error
	log.Fatal(err)
}
