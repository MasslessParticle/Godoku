package main

import (
	"encoding/json"
	"fmt"
	"github.com/masslessparticle/sudokusolver/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/masslessparticle/sudokusolver/db"
	"github.com/masslessparticle/sudokusolver/domain"
	"github.com/masslessparticle/sudokusolver/sudoku"
	"net/http"
	"strconv"
)

func SavePuzzleHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		puzzle := domain.Puzzle{Content: request.FormValue("puzzle")}
		puzzleId := strconv.Itoa(db.InsertPuzzle(puzzle))

		response.Write([]byte(puzzleId))
	}
}

func GetPuzzleHandler(solved bool) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["id"]
		puzzle := db.GetPuzzle(id)

		if solved {
			solved := sudoku.Solve(puzzle.Content)
			puzzle.Content = fmt.Sprintf(solved.String())
		}

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
