package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-pg/pg"
)

// Shared instance of database connection
var db *pg.DB

// Internal errors
var errDBNotInitiated = errors.New("db: not initiated")

// handleConnectError checks if the error
// contains "connection refuse" or "no such host"
// then waits 3 seconds to retry connecting afterwards
func handleConnectError(connectError error) {
	connectRefuse := "5432: connect: connection refused"
	noSuch := "no such host"
	containsRefused := strings.Contains(connectError.Error(), connectRefuse)
	containsNoSuch := strings.Contains(connectError.Error(), noSuch)

	if containsRefused || containsNoSuch {
		log.Println("db not ready yet!")
		threeSeconds := time.Duration(3) * time.Second
		time.Sleep(threeSeconds)
	}
}

// checkDB checks the database connection
// by performing a simple SELECT query
func checkDB() error {
	if db == nil {
		return errDBNotInitiated
	}

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

// connectToDB first checks if there's a valid
// DB connection returns it, otherwise tries to
// connect/reconnect to the DB
func connectToDB() *pg.DB {
	for {
		if err := checkDB(); err == nil {
			return db
		} else if err != errDBNotInitiated {
			handleConnectError(err)
		}

		hostPortString := fmt.Sprintf("%s:%s", os.Getenv("PGHOST"), os.Getenv("PGPORT"))

		options := &pg.Options{
			User:     os.Getenv("PGUSER"),
			Password: os.Getenv("PGPASSWORD"),
			Database: os.Getenv("PGDATABASE"),
			Addr:     hostPortString,
		}

		db = pg.Connect(options)
	}
}
