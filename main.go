package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

// config model
type config struct {
	addr    string `default:":4000"`
	name    string `default:"Simple Go Microservice 🚀"`
	version string `default:"0.0.1"`
}

// methods
func (cnf *config) init() {
	var addr, name, vers string

	addr = os.Getenv("ADDR")
	if addr != "" {
		cnf.addr = addr
	}

	name = os.Getenv("NAME")
	if name != "" {
		cnf.name = name
	}

	vers = os.Getenv("VERS")
	if vers != "" {
		cnf.version = vers
	}
}

// handler
func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Server", "Simple Go Service")
	w.Write([]byte(`{"message":"Hi there 👋"}`))
}

func help(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Server", "Simple Go Service")

	w.Write(body)
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

// execution entry point
func main() {
	// declare config
	var cfg config

	// init config
	cfg.init()

	// init logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// create a new servermux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("GET /{$}", hello)
	mux.HandleFunc("GET /help", help)
	mux.HandleFunc("GET /{user}/data/{datatype}", viewData)
	mux.HandleFunc("POST /{user}/data/{datatype}", uploadData)

	// log the service startup
	logger.Info(fmt.Sprintf("Starting service %s v:%s  on %s", cfg.name, cfg.version, cfg.addr))

	// start http server on :4000
	err := http.ListenAndServe(cfg.addr, mux)
	// log the error message if ListenAndServe() encounters error
	logger.Error(err.Error())
	os.Exit(1)
}
