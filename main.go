package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rikkinovar/movie-catalog-api/db"
)

func main() {
	godotenv.Load()

	dbConn, err := db.GetConn()
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{}

	initRepositories(dbConn, httpClient)
	initServices()

	switch command() {
	case "serve":
		serveHTTP()
	case "migrate":
		db.Migrate(dbConn)
	default:
		fmt.Println("Invalid command")
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
