package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then register the home function as the handler for the "/" URL path.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) // The {$} prevents the home function from being called for any URL path that starts with "/".
	mux.HandleFunc("GET /character/view/{id}", characterView)
	mux.HandleFunc("GET /character/create", characterCreate)
	mux.HandleFunc("POST /character/create", characterCreatePost)

	// Print a log message indicating that the server is starting on port 4000.
	log.Print("Starting server on port 4000")

	// Use the http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
