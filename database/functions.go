package database

import (
	"log"
	"mphclub-rest-server/models"
	"time"

	"github.com/rs/xid"
)

func CreateUser(u models.UserInfo) error {
	db := connectToDB()

	err := db.Insert(&u)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("user created")
	return nil
}

func UpsertListing(v models.Vehicle) error {
	db := connectToDB()

	carParam := &v
	err := db.Select(carParam)
	if err != nil {
		log.Println("car does not exist, create")
	} else {
		log.Println("car does exist, update")
		return nil
	}

	v.ID = xid.New().String()
	v.CreatedTime = time.Now()

	log.Println(v)

	err = db.Insert(&v)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(v)
	log.Println("vehicle created")
	return nil
}

func GetCars() ([]models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	err := db.Model(&vehicleList).Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return vehicleList, nil
}
