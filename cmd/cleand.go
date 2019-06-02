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
	r.Handle("/v1/tags", handlers.CompressHandler(&HtmlCleaner.CleanerPostBody{cln})).Methods("POST")
	r.Handle("/v1/regex", handlers.CompressHandler(&HtmlCleaner.CleanerPostTitle{cln})).Methods("POST")
	r.Handle("/v1/links", handlers.CompressHandler(&HtmlCleaner.LinkifyPost{cln})).Methods("POST")
	r.Handle("/v1/all", handlers.CompressHandler(&HtmlCleaner.AllPost{cln})).Methods("POST")

	// Bind to a port and pass our router in
	log.Print("starting server...")
	server := &http.Server{
		Addr:    ":8002",
		Handler: r,
	}
	log.Print(server.ListenAndServe())
}
