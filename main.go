package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there 👋\n"))
}

func viewData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display user data 📁\n"))
}

func uploadData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload user data 📤\n"))
}

func main() {
	// create a new servermux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/{$}", hello)
	mux.HandleFunc("/data/view", viewData)
	mux.HandleFunc("/data/upload", uploadData)

	// log the service startup
	log.Print("starting service on :4000")

	// start http server on :4000
	err := http.ListenAndServe(":4000", mux)
	// log the error message if ListenAndServe() encounters error
	log.Fatal(err)
}
