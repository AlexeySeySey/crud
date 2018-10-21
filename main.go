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

	router.HandleFunc("/new", notes.NewOne).Methods("POST")
	router.HandleFunc("/", servs.Main).Methods("GET")
	router.HandleFunc("/fetch", notes.FetchAll).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./templates/src")))

	fmt.Print("Server running ")

	log.Fatal(http.ListenAndServe(":8000", router))

}
