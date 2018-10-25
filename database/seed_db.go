package database

import (
	"github.com/go-pg/pg"
)

func seedDB() {
	db := connectToDB()

	seedVehicles(db)
}

func seedVehicles(db *pg.DB) {
}
