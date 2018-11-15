package database

import (
	"errors"
	"log"
	"mphclub-rest-server/models"
	"time"

	"github.com/go-pg/pg/orm"
	"github.com/rs/xid"
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

func EditPhotoURLArrayOnVehicle(vehicleID, photoURL string) error {
	db := connectToDB()
	vehicleToAttach := &models.Vehicle{
		ID: vehicleID,
	}

	err := db.Select(vehicleToAttach)
	if err != nil {
		log.Println(err)
		return err
	}

	vehicleToAttach.Photos = append(vehicleToAttach.Photos, photoURL)
	vehicleToAttach.UpdatedTime = time.Now()

	_, err = db.Model(vehicleToAttach).Column("photos", "updated_time").WherePK().Update()
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

		v = v.Merge(car)
		v.UpdatedTime = time.Now()

		if dbErr := db.Update(&v); dbErr != nil {
			return "", "", dbErr
		}
		return v.ID, "updated", nil
	}
	v.ViewIndex = -1
	v.ID = xid.New().String()
	v.CreatedTime = time.Now()
	v.Status = "PENDING"

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

func GetMyCars(u *models.User) ([]models.Vehicle, error) {
	db := connectToDB()

	var users []models.User

	err := db.Model(&users).
		Column("user.*", "Vehicles").
		Relation("Vehicles", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Where("id = ?", u.ID).
		Select()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(users[0].Vehicles) == 0 {
		return nil, errors.New("This user has no cars")
	}

	return users[0].Vehicles, nil
}

func GetExplore() ([][]models.Vehicle, error) {
	var explore3 [][]models.Vehicle
	var vehicle models.Vehicle
	var carTypes []string

	db := connectToDB()

	if err := db.
		Model(&vehicle).
		ColumnExpr("DISTINCT vehicle.vehicle_type").
		Where("vehicle_type IS NOT NULL").
		Select(&carTypes); err != nil {
		return nil, err
	}

	for _, carType := range carTypes {
		var list []models.Vehicle
		var err error

		list, err = getTypeVehicleArray(carType)
		if err != nil {
			return nil, err
		}

		explore3 = append(explore3, list)
	}

	return explore3, nil
}

func getTypeVehicleArray(carType string) ([]models.Vehicle, error) {
	var list []models.Vehicle

	db := connectToDB()

	if err := db.
		Model(&list).
		Where("vehicle_type = ?", carType).
		Limit(3).
		Select(); err != nil {
		return nil, err
	}

	return list, nil
}
