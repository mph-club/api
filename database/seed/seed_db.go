package seed

import (
	"github.com/go-pg/pg"
)

// CheckForSeed checks all models seeds
// and run them if they're not applied before
func CheckForSeed(db *pg.DB) {
	checkForVehiclesSeed(db)
}
