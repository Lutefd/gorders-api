package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/hello", hello)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
