package main

import (
	"github.com/gorilla/mux"
	"log"
	"ms-greeter/handler"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", handler.Health())
	r.HandleFunc("/readiness", handler.Readiness())
	r.HandleFunc("/greet", handler.Greet())
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
