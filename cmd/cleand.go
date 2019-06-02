package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	HtmlCleaner "github.com/untoreh/cleanup"
)

func main() {
	log.Print("loading cleaner...")
	cln := HtmlCleaner.New()
	r := mux.NewRouter()
	// r.Handle("/", handlers.CompressHandler().Methods("GET")
	r.Handle("/v1/body", handlers.CompressHandler(&HtmlCleaner.CleanerPostBody{cln})).Methods("POST")
	r.Handle("/v1/title", handlers.CompressHandler(&HtmlCleaner.CleanerPostTitle{cln})).Methods("POST")

	// Bind to a port and pass our router in
	log.Print("starting server...")
	server := &http.Server{
		Addr:    ":8002",
		Handler: r,
	}
	log.Print(server.ListenAndServe())
}
