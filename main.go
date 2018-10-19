package main

/*
-Make input for sql-connection parametrs.Open params for any specific connection.
*/

import (
	"./servs"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"./monkey/gorilla/mux"
)

var notes = servs.Note{}

func Main(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func NewOne(w http.ResponseWriter, r *http.Request) {
	// ! putting data to database and return answer

	// .. Continue here: you can't get data from post ajax by axios, it's [empty]

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", Main).Methods("GET")
	router.HandleFunc("/new", NewOne).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./templates/src")))

	fmt.Print("Server running ")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Print(".")
			time.Sleep(time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(":8000", router))

	wg.Wait()

}
