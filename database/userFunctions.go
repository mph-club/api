package database

import (
	"log"
	"mphclub/mphclub-rest-server/models"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
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
		Column("user.*").
		Relation("DriverLicense").
		Relation("Insurance").
		Where("\"user\".id = ?", userID).
		Select(); err != nil {
		return models.User{}, err
	}

	users[0].InsuranceMap = echo.Map{
		"insurance_name": users[0].Insurance.InsuranceName,
		"policy_number":  users[0].Insurance.PolicyNumber,
	}

	return users[0], nil
}

func AddUserPhotoURL(userID, photoURL string) error {
	db = connectToDB()
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
	db = connectToDB()
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

func GetHostDetails(host models.User) (models.User, error) {
	db := connectToDB()

	var users []models.User

	err := db.Model(&users).
		Column("user.*", "Vehicles").
		Relation("Vehicles", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Column("user.*", "DriverLicense").
		Relation("DriverLicense").
		Where("\"user\".id = ?", host.ID).
		Select()

	if err != nil {
		return models.User{}, err
	}

	return users[0], nil
}

func GetMyReservations(renterID string) ([]models.Trip, error) {
	db = connectToDB()

	var trips []models.Trip

	err := db.Model(&trips).
		Column("trip.*").
		Relation("User").
		Relation("Vehicle").
		Where("\"user\".id = ?", renterID).
		Select()

	if err != nil {
		return nil, err
	}

	return trips, nil
}

func AddInsurance(userInsurance models.Insurance, userID string) error {
	db = connectToDB()

	if err := db.Insert(&userInsurance); err != nil {
		return err
	}

	user := models.User{
		ID:          userID,
		InsuranceID: userInsurance.ID,
	}

	_, err := db.Model(&user).
		Column("insurance_id").
		WherePK().
		Update()

	if err != nil {
		return err
	}

	return nil
}
