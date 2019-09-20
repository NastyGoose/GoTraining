package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/commands", CommandsHandler).Methods(http.MethodPost)

	log.Println("Server is up!")

	http.ListenAndServe(":8080", r)
}
