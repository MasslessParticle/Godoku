package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.
	HandleFunc("/puzzle", SavePuzzleHandler()).
	Methods("POST")
	router.
	HandleFunc("/puzzle/{id}", GetPuzzleHandler()).
	Methods("GET")
	router.
	HandleFunc("/solved/{id}", GetPuzzleHandler()).
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