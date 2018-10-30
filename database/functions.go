package database

import (
	"log"
	"mphclub-rest-server/models"

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

func CreateListing(v models.Vehicle) error {
	db := connectToDB()

	v.ID = xid.New().String()

	err := db.Insert(&v)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("vehicle created")
	return nil
}
