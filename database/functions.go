package database

import (
	"log"
	"mphclub-rest-server/models"
)

func CreateUser(u models.UserInfo) error {
	db := connectToDB()

	err := db.Insert(&u)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
