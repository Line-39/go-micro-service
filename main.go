package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there 👋\n"))
}

func viewData(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	if user == "" {
		http.NotFound(w, r)
		return
	}
	dtype := r.PathValue("datatype")
	if dtype == "" {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("📁 Display the %s data for user %s\n", dtype, user)
	w.Write([]byte(msg))

}

func uploadData(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	if user == "" {
		http.NotFound(w, r)
		return
	}
	dtype := r.PathValue("datatype")
	if dtype == "" {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("📤 Upload the %s data for user %s\n", dtype, user)
	w.Write([]byte(msg))
}

func main() {
	// create a new servermux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/{$}", hello)
	mux.HandleFunc("/{user}/data/{datatype}/view", viewData)
	mux.HandleFunc("/data/upload", uploadData)

	// log the service startup
	log.Print("starting service on :4000")

	// start http server on :4000
	err := http.ListenAndServe(":4000", mux)
	// log the error message if ListenAndServe() encounters error
	log.Fatal(err)
}
