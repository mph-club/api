package database

import (
	"fmt"
	"log"
	pb "mphclub-server/api-generated"

	"github.com/go-pg/pg"
)

func seedDB() {
	db := connectToDB()

	seedVehicles(db)
}

func seedVehicles(db *pg.DB) {
	var vehicles []*pb.Vehicle
	err := db.Model(&vehicles).Select()
	if err != nil {
		log.Println("something went wrong: ", err)
	}
	if len(vehicles) == 0 {

		someVehicle := &pb.Vehicle{
			Color:       "yellow",
			DayMax:      10,
			DayMin:      1,
			Description: "This is a car",
			Make:        "chevy",
			Model:       "cavalier",
			Seats:       "3",
			Status:      "not reserved",
			Trim:        "tan",
			VehicleType: "sedan",
			Year:        "2008",
		}

		for i := 0; i < 31; i++ {
			vehicles = append(vehicles, someVehicle)
		}

		for i, vehicle := range vehicles {
			var idString = fmt.Sprintf("%.3d", i)
			vehicle.Id = idString

			err := db.Insert(vehicle)
			if err != nil {
				log.Println("error!")
				log.Println(err)
			}
		}
	} else {
		log.Println("already seeded")
	}

}
