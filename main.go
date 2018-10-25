package main

import (
	"fmt"
	"log"
	"net/http"

	"./monkey/gorilla/mux"
	"./servs"
)

var notes = servs.Note{}
var router = mux.NewRouter()

func main() {

	router.HandleFunc("/new", servs.NewOne).Methods("POST")
	router.HandleFunc("/", servs.Main).Methods("GET")
	router.HandleFunc("/fetch", notes.FetchAll).Methods("GET")
	router.HandleFunc("/delete/{id}", servs.DropOne).Methods("DELETE")
	router.HandleFunc("/update/{id}", servs.UpdateOne).Methods("PUT")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./templates/src")))

	fmt.Print("Server running... ")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}

}
