package main

import (
	"fmt"
	"github.com/masslessparticle/sudokusolver/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/masslessparticle/sudokusolver/db"
	"log"
	"net/http"
	"os"
)

func main() {
	db.InitDB()

	router := mux.NewRouter()

	router.
		HandleFunc("/puzzle", SavePuzzleHandler()).
		Methods("POST")
	router.
		HandleFunc("/puzzle/{id}", GetPuzzleHandler(false)).
		Methods("GET")
	router.
		HandleFunc("/solved/{id}", GetPuzzleHandler(true)).
		Methods("GET")
	router.
		HandleFunc("/", BasicResponseHandler()).
		Methods("GET")

	if err := http.ListenAndServe(fmt.Sprintf(":%v", getPort()), router); err != nil {
		log.Fatalln(err)
	}
}

func getPort() string {
	if configuredPort := os.Getenv("PORT"); configuredPort == "" {
		return "3000"
	} else {
		return configuredPort
	}
}
