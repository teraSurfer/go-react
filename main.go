package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	// "services"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", status).Methods("GET")

	log.Fatal(http.ListenAndServe(":3030", router))
}

func main() {
	fmt.Println("Starting server on port :3030")

	// InitialMigration()

	handleRequests()
}
