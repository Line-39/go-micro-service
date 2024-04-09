package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Server", "Simple Go Service")
	w.Write([]byte(`{"message":"Hi there 👋"}`))
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

	w.WriteHeader(http.StatusCreated)

	msg := fmt.Sprintf("📤 Upload the %s data for user %s\n", dtype, user)

	w.Write([]byte(msg))
}

func main() {
	// create a new servermux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("GET /{$}", hello)
	mux.HandleFunc("GET /{user}/data/{datatype}", viewData)
	mux.HandleFunc("POST /{user}/data/{datatype}", uploadData)

	// log the service startup
	log.Print("starting service on :4000")

	// start http server on :4000
	err := http.ListenAndServe(":4000", mux)
	// log the error message if ListenAndServe() encounters error
	log.Fatal(err)
}
