package database

import (
	"fmt"

	"github.com/go-pg/pg"
)

func seedDB() {
	db := connectToDB()

	seedVehicles(db)
}

func seedVehicles(db *pg.DB) {
	fmt.Println("seeded")
}
