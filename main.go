package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Go Webserver 🚀")
}

// Handler for "/data"
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Example JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status": "ok", "message": "Hello from Go"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/data", dataHandler)

	server := &http.Server{
		Addr:    ":8080", // runs on http://localhost:8080
		Handler: mux,
	}

	log.Println("Server started on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
