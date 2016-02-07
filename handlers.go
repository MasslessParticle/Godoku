package main

import (
	"net/http"
	"github.com/masslessparticle/sudokusolver/domain"
	"encoding/json"
	"github.com/masslessparticle/sudokusolver/db"
	"strconv"
	"github.com/gorilla/mux"
)

func SavePuzzleHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		puzzle := domain.Puzzle{Content: request.FormValue("puzzle")}
		puzzleId := strconv.Itoa(db.InsertPuzzle(puzzle))

		response.Write([]byte(puzzleId))
	}
}

func GetPuzzleHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["id"]
		puzzle := db.GetPuzzle(id)

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