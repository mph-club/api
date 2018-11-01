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

func UpsertListing(v models.Vehicle) (string, string, error) {
	db := connectToDB()

	car := &v
	if err := db.Select(car); err != nil {
		log.Println("car does not exist, create")
	} else {
		log.Println("car does exist, update")
		log.Println(v)
		car.UpdatedTime = time.Now()

		if dbErr := db.Update(car); dbErr != nil {
			return "", "", dbErr
		}
		return car.ID, "updated", nil
	}
	car.ID = xid.New().String()
	car.CreatedTime = time.Now()

	if err := db.Insert(car); err != nil {
		log.Println(err)
		return "", "", err
	}

	log.Println("vehicle created")
	return car.ID, "created", nil
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
