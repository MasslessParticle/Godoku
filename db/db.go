package db
import (
	"os"
	"gopkg.in/gorp.v1"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masslessparticle/sudokusolver/domain"
	"net/url"
	"fmt"
	"strconv"
)

var dbmap *gorp.DbMap

func InitDB() {
	db, err := sql.Open("mysql", databaseUrl())
	checkError(err, "Failed to establish database connection")

	dbmap = &gorp.DbMap{Db: db, Dialect:gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(domain.Puzzle{}, "puzzle").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkError(err, "Table Creation Failed")
}

func InsertPuzzle(puzzle domain.Puzzle) int {
	dbmap.Insert(&puzzle)
	return puzzle.Id
}

func GetPuzzle(idString string) *domain.Puzzle {
	id, err := strconv.Atoi(idString)
	checkError(err, "Error Determining Puzzle ID")

	obj, err := dbmap.Get(domain.Puzzle{}, id)
	checkError(err, "Error Loading Puzzle")

	return obj.(*domain.Puzzle)
}

func databaseUrl() string {
	urlPath := os.Getenv("DATABASE_URL")
	if urlPath == "" {
		urlPath = "mysql2://sudoku:sudoku@127.0.0.1:3306/sudokusolver_development?reconnect=true"
	}

	databaseUrl, err := url.Parse(urlPath)
	checkError(err, "Error parsing DATABASE_URL")

	return formattedUrl(databaseUrl)
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func formattedUrl(url *url.URL) string {
	return fmt.Sprintf(
		"%v@tcp(%v)%v?parseTime=true",
		url.User,
		url.Host,
		url.Path,
	)
}
