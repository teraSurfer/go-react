package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	// folder path.
	"github.com/teraSurfer/go-react/services"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", status).Methods("GET")
	router.HandleFunc("/movies", services.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies", services.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", services.GetMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", services.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", services.DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3030", router))
}

func main() {
	fmt.Println("Starting server on port :3030")

	services.InitialMigration()

	handleRequests()
}
