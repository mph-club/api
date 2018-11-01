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

	car := models.Vehicle{
		ID: v.ID,
	}

	if err := db.Select(&car); err != nil {
		log.Println(err.Error())
		log.Println("car does not exist, create")
	} else {
		log.Println("car does exist, update")

		log.Println(v.Merge(car))
		v = v.Merge(car)

		v.UpdatedTime = time.Now()
		log.Println(&v)

		if dbErr := db.Update(&v); dbErr != nil {
			return "", "", dbErr
		}
		return v.ID, "updated", nil
	}

	v.ID = xid.New().String()
	v.CreatedTime = time.Now()

	if err := db.Insert(&v); err != nil {
		log.Println(err)
		return "", "", err
	}

	log.Println("vehicle created")
	return v.ID, "created", nil
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
