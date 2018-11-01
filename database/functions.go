package database

import (
	"log"
	"mphclub-rest-server/models"
	"time"

	"github.com/rs/xid"
)

func UpsertUser(u models.UserInfo) error {
	db := connectToDB()

	user := models.UserInfo{
		Sub: u.Sub,
	}

	if err := db.Select(&user); err != nil {
		log.Println(err.Error())
		log.Println("user does not exist, create")
	} else {
		log.Println("user does exist, update")

		u = u.Merge(user)

		if dbErr := db.Update(&u); dbErr != nil {
			return dbErr
		}
		return err
	}

	if err := db.Insert(&user); err != nil {
		return err
	}

	return nil
}

func EditPhotoURLArrayOnVehicle(vehicleID, photoURL string) error {
	db := connectToDB()
	vehicleToAttach := &models.Vehicle{ID: vehicleID}

	err := db.Select(vehicleToAttach)
	if err != nil {
		log.Println(err)
		return err
	}

	vehicleToAttach.Photos = append(vehicleToAttach.Photos, photoURL)
	_, err = db.Model(vehicleToAttach).Column("photos").WherePK().Update()
	if err != nil {
		log.Println(err)
		return err
	}

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
