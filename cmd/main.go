package main

import (
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/getMemoryData/{key}", handlers.GetMemoryData).Methods(http.MethodGet)
	router.HandleFunc("/setMemoryData", handlers.SetMemoryData).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
