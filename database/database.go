package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"mphclub-rest-server/database/seed"
	"mphclub-rest-server/models"

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

// createTypes attempts to create
// all custom-defined types
func createTypes(db *pg.DB) error {
	types := map[string]string{
		"status":       "ENUM ('APPROVED', 'DENIED', 'PENDING')",
		"transmission": "ENUM ('AUTO', 'MANUAL')",
		"miles":        "ENUM ('0-50', '50-100', '100-130', '130+')",
	}

	for name, t := range types {
		q := fmt.Sprintf("CREATE TYPE %v AS %v", name, t)
		if _, err := db.Exec(q); err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("type \"%v\" already exists", name)
			} else {
				return err
			}
		} else {
			log.Printf("type \"%v\" created successfully", name)
		}
	}

	return nil
}

// CreateSchema attempts to create
// all existing schemas
func CreateSchema() {
	db := connectToDB()

	if err := createTypes(db); err != nil {
		log.Panicln(err)
	}

	for _, model := range []interface{}{
		&models.Vehicle{},
		&models.User{},
		&models.UserNote{},
		&models.VehicleNote{},
		&models.DriverLicense{},
		&models.Staff{},
		//models go here
	} {
		name := reflect.TypeOf(model).Elem().Name()

		if err := db.CreateTable(model, nil); err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("schema \"%v\" already exists", name)
			} else {
				log.Panicln(err)
			}
		} else {
			log.Printf("schema \"%v\" created successfully", name)
		}
	}

	seed.CheckForSeed(db)
}
