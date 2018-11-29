package database

import (
	"log"
	"mphclub-rest-server/models"

	"github.com/go-pg/pg"
	"github.com/rs/xid"
)

func createEnums(db *pg.DB) {
	qs := []string{
		"CREATE TYPE status AS ENUM ('APPROVED', 'DENIED', 'PENDING')",
		"CREATE TYPE transmission AS ENUM ('AUTO', 'MANUAL')",
		"CREATE TYPE miles AS ENUM ('0-50', '50-100', '100-130', '130+')",
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	CreateSchema()
}

func seedDB() {
	db := connectToDB()

	checkForSeed(db)
}

func seedVehicles(db *pg.DB) {
	err := db.Insert(&carList)
	if err != nil {
		log.Println(err)
	}
}

func checkForSeed(db *pg.DB) {
	var lambo models.Vehicle

	err := db.
		Model(&lambo).
		Where("license_plate = ?", "DQQ R63").
		Select()

	if err != nil {
		seedVehicles(db)
	} else {
		log.Println("already seeded")
	}
}

var carList = []models.Vehicle{
	{
		Make:         "Lamborghini",
		Model:        "Aventador Roadster",
		Year:         2014,
		Trim:         "",
		Color:        "Dark Grey",
		Doors:        2,
		Seats:        2,
		Description:  "lamborghini aventador roadster",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "DQQ R63",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Lamborghini",
		Model:        "Aventador",
		Year:         2015,
		Trim:         "",
		Color:        "White",
		Doors:        2,
		Seats:        2,
		Description:  "white lamborghini aventador",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "DQQ T73",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Lamborghini",
		Model:        "Gallardo Spyder",
		Year:         2016,
		Trim:         "",
		Color:        "White",
		Doors:        2,
		Seats:        2,
		Description:  "convertible white lamborghini gallardo spyder",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Lamborghini",
		Model:        "Huracan",
		Year:         2015,
		Trim:         "",
		Color:        "Dark Grey",
		Doors:        2,
		Seats:        2,
		Description:  "white lamborghini huracan",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "ECM G46",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Lamborghini",
		Model:        "Huracan Spyder",
		Year:         2016,
		Trim:         "",
		Color:        "White",
		Doors:        2,
		Seats:        2,
		Description:  "convertible white lamborghini huracan spyder",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "KEBY31",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Ferrari",
		Model:        "458 Spider",
		Year:         2015,
		Trim:         "",
		Color:        "Blue",
		Doors:        2,
		Seats:        2,
		Description:  "convertible blue ferrari 458 spider",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "DSD F25",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Ferrari",
		Model:        "458 Spider",
		Year:         2016,
		Trim:         "",
		Color:        "yellow",
		Doors:        2,
		Seats:        2,
		Description:  "convertible yellow ferrari 458 spider",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Ferrari",
		Model:        "California T",
		Year:         2016,
		Trim:         "",
		Color:        "Red",
		Doors:        2,
		Seats:        4,
		Description:  "convertible red ferrari california T",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "CDZ 574",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Ferrari",
		Model:        "488 GTB",
		Year:         2016,
		Trim:         "",
		Color:        "light Grey",
		Doors:        2,
		Seats:        2,
		Description:  "grey ferrari 488 GTB",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "DQQ W25",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Audi",
		Model:        "R8 V10 Plus",
		Year:         2016,
		Trim:         "",
		Color:        "light grey",
		Doors:        2,
		Seats:        2,
		Description:  "Amazing Audi R8 V10 Plus",
		VehicleType:  "Sports",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "BMW",
		Model:        "i8",
		Year:         2015,
		Trim:         "",
		Color:        "dark grey",
		Doors:        2,
		Seats:        4,
		Description:  "magnificent BMW i8",
		VehicleType:  "Sport",
		Miles:        "0-50",
		LicensePlate: "EQJ S46",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Bentley",
		Model:        "GTC",
		Year:         2016,
		Trim:         "",
		Color:        "white",
		Doors:        2,
		Seats:        2,
		Description:  "Convertible white bentley gtc",
		VehicleType:  "Sedan",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Rolls Royce",
		Model:        "Ghost",
		Year:         2016,
		Trim:         "",
		Color:        "black",
		Doors:        4,
		Seats:        5,
		Description:  "luxurious black rolls royce ghost",
		VehicleType:  "Sedan",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Rolls Royce",
		Model:        "Ghost Series 2",
		Year:         2016,
		Trim:         "",
		Color:        "white",
		Doors:        4,
		Seats:        5,
		Description:  "luxurious white rolls royce ghost series 2",
		VehicleType:  "Sedan",
		Miles:        "0-50",
		LicensePlate: "CMH Y38",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Rolls Royce",
		Model:        "Phantom Drophead",
		Year:         2016,
		Trim:         "",
		Color:        "black",
		Doors:        4,
		Seats:        5,
		Description:  "luxurious black rolls royce phantom drophead",
		VehicleType:  "Sedan",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Rolls Royce",
		Model:        "Wraith",
		Year:         2016,
		Trim:         "",
		Color:        "white",
		Doors:        2,
		Seats:        4,
		Description:  "luxurious white rolls royce wraith",
		VehicleType:  "Sedan",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Rolls Royce",
		Model:        "Dawn",
		Year:         2016,
		Trim:         "",
		Color:        "white",
		Doors:        2,
		Seats:        4,
		Description:  "luxurious white rolls royce dawn",
		VehicleType:  "Sedan",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Range Rover",
		Model:        "Startech HSE",
		Year:         2016,
		Trim:         "HSE",
		Color:        "white",
		Doors:        4,
		Seats:        5,
		Description:  "white startech range rover HSE",
		VehicleType:  "SUV",
		Miles:        "0-50",
		LicensePlate: "",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Range Rover",
		Model:        "Supercharged LWB",
		Year:         2016,
		Trim:         "LWB",
		Color:        "black",
		Doors:        4,
		Seats:        5,
		Description:  "black range rover Supercharged LWB",
		VehicleType:  "SUV",
		Miles:        "0-50",
		LicensePlate: "GQP M15",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
	{
		Make:         "Mercedes Benz",
		Model:        "G63 AMG",
		Year:         2016,
		Trim:         "AMG",
		Color:        "black",
		Doors:        4,
		Seats:        5,
		Description:  "black Mercedes Benz G63 AMG",
		VehicleType:  "SUV",
		Miles:        "0-50",
		LicensePlate: "CHA9637",
		Transmission: "AUTO",
		ID:           xid.New().String(),
		Status:       "PENDING",
	},
}
