package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Answer struct {
	Message string `json:"message"`
}

func Greet() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		answer := Answer{
			Message: "Hi stranger!",
		}

		log.Print("greeting request, message: ", answer.Message)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(answer)
	}
}
