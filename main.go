package main

import (
	"github.com/gorilla/mux"
	"time"
	"net/http"
	"log"
	"github.com/sansaralab/buddha/handlers"
)

func main() {
	r := mux.NewRouter()
	setRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func setRoutes(r *mux.Router) {
	r.Handle("/projects", handlers.NewProjectHandler())
}
