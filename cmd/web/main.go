package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Use slog.New() to initialize a new logger with some additional configuration.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Use the http.NewServeMux() function to initialize a new servemux, then register the home function as the handler for the "/" URL path.
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", home) // The {$} prevents the home function from being called for any URL path that starts with "/".
	mux.HandleFunc("GET /character/view/{id}", characterView)
	mux.HandleFunc("GET /character/create", characterCreate)
	mux.HandleFunc("POST /character/create", characterCreatePost)

	// Print a log message indicating that the server is starting on port 4000.
	logger.Info("Starting server", "addr", *addr)

	// Use the http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
