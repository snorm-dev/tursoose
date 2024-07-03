package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Could not load .env file: ", err)
	}

	if len(os.Args) == 0 {
		log.Fatal("error: no goose command given. Use goose's help for more info")
	}

	command := os.Args[1]

	args := os.Args[2:]

	dir, ok := os.LookupEnv("TURSOOSE_DIR")
	if !ok {
		log.Fatal("error: no schema directory")
	}

	url, ok := os.LookupEnv("TURSOOSE_URL")
	if !ok {
		log.Fatal("error: no database url")
	}

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal("error setting dialect:", err)
	}

	if err := goose.RunContext(context.Background(), command, db, dir, args...); err != nil {
		log.Fatal("error running goose:", err)
	}
}
