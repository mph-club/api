package database

import (
	"log"
	"mphclub-rest-server/models"
)

func UpsertUser(u models.User) error {
	db := connectToDB()

	user := models.User{
		ID: u.ID,
	}

	if err := db.Select(&user); err != nil {
		log.Println(err.Error())
		log.Println("user does not exist, create")

		user = user.Merge(u)
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

func GetUser(userID string) (models.User, error) {
	var users []models.User
	db := connectToDB()

	if err := db.Model(&users).
		Column("user.*", "DriverLicense").
		Relation("DriverLicense").
		Where("\"user\".id = ?", userID).
		Select(); err != nil {
		return models.User{}, err
	}

	return users[0], nil
}

func AddUserPhotoURL(userID, photoURL string) error {
	db := connectToDB()
	user := &models.User{
		ID: userID,
	}
	err := db.Select(user)
	if err != nil {
		return err
	}

	user.ProfilePhotoURL = photoURL
	_, err = db.Model(user).Column("profile_photo_url").WherePK().Update()
	if err != nil {
		return err
	}

	return nil
}

func AddDriverLicense(userID string, dl *models.DriverLicense) error {
	db := connectToDB()
	if err := db.Insert(dl); err != nil {
		return err
	}

	user := models.User{
		ID:              userID,
		DriverLicenseID: dl.ID,
	}

	_, err := db.Model(&user).
		Column("driver_license_id").
		WherePK().
		Update()

	if err != nil {
		return err
	}

	return nil
}

func GetDriverLicense(userID string) (models.DriverLicense, error) {
	db := connectToDB()
	var u []models.User

	if err := db.Model(&u).
		Column("user.*", "DriverLicense").
		Select(); err != nil {
		return models.DriverLicense{}, err
	}

	return u[0].DriverLicense, nil
}

func EditOfacStatus(userID string, ofacCheck bool) error {
	user := models.User{
		ID:        userID,
		OfacCheck: ofacCheck,
	}

	db := connectToDB()

	_, err := db.Model(&user).
		Column("ofac_check").
		WherePK().
		Update()

	if err != nil {
		return err
	}

	return nil
}