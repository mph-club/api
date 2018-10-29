package database

import (
	"fmt"
	"log"

	"mphclub-rest-server/models"

	"github.com/go-pg/pg"
)

func seedDB() {
	db := connectToDB()

	seedVehicles(db)
}

func seedVehicles(db *pg.DB) {
	var vehicles []*models.Vehicle
	err := db.Model(&vehicles).Select()
	if err != nil {
		log.Println("something went wrong: ", err)
	}
	fmt.Println(vehicles)

	if len(vehicles) == 0 {

		someVehicle := &models.Vehicle{
			Color:       "yellow",
			DayMax:      10,
			DayMin:      1,
			Description: "This is a car",
			Make:        "chevy",
			Model:       "cavalier",
			Seats:       3,
			Status:      "not reserved",
			Trim:        "tan",
			VehicleType: "sedan",
			Year:        2008,
		}

		for i := 0; i < 31; i++ {
			vehicles = append(vehicles, someVehicle)
		}

		for i, vehicle := range vehicles {
			var idString = fmt.Sprintf("%.3d", i)
			vehicle.ID = idString

			err := db.Insert(vehicle)
			if err != nil {
				log.Println("error!")
				log.Println(err)
			}
		}
	} else {
		fmt.Println("already seeded")
	}
}
