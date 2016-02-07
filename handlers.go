package main

import (
	"net/http"
	"github.com/masslessparticle/sudokusolver/domain"
	"encoding/json"
)

func SavePuzzleHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		puzzle := domain.Puzzle{Id: 6}

		js, err := json.Marshal(puzzle)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	}
}

func GetPuzzleHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		puzzle := domain.Puzzle{Id: 6, Content: "This is where the puzzle will go"}

		js, err := json.Marshal(puzzle)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	}
}

func BasicResponseHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("The Solver is responding"))
	}
}