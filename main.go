package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	the_server := &http.Server{
		Addr:    ":" + "8080",
		Handler: r,
	}

	log.Fatal(the_server.ListenAndServe())
}
