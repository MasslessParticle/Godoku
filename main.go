package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"os"
	"github.com/masslessparticle/sudokusolver/db"
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