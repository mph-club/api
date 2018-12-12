package database

import (
	"fmt"
	"log"
	"mphclub-rest-server/models"
	"net/url"
	"strings"
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

func EditPhotoURLArrayOnVehicle(vehicleID, filename string) error {
	db := connectToDB()

	url := fmt.Sprintf("https://mphclub.s3.amazonaws.com/%s/%s", vehicleID, filename)
	thumbnailURL := fmt.Sprintf("https://mphclub.s3.amazonaws.com/%s/thumb/%s", vehicleID, filename)

	vehicleToAttach := &models.Vehicle{
		ID: vehicleID,
	}

	err := db.Select(vehicleToAttach)
	if err != nil {
		log.Println(err)
		return err
	}

	vehicleToAttach.Photos = append(vehicleToAttach.Photos, url)
	vehicleToAttach.Thumbnails = append(vehicleToAttach.Thumbnails, thumbnailURL)
	vehicleToAttach.UpdatedTime = time.Now()

	_, err = db.Model(vehicleToAttach).
		Column("photos", "updated_time", "thumbnails").
		WherePK().
		Update()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func AddUserPhotoURL(userID, photoURL string) error {
	db := connectToDB()
	user := &models.User{
		ID: userID,
	}
	err := db.Select(user)
	if err != nil {
		log.Println(err)
		return err
	}

	user.ProfilePhotoURL = photoURL
	_, err = db.Model(user).Column("profile_photo_url").WherePK().Update()
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

	if v.ViewIndex == 0 {
		v.ViewIndex = -1
	}
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

func GetCars(queryParams url.Values, carType string) (int, []models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	if len(carType) == 0 {
		count, err := db.Model(&vehicleList).
			Apply(orm.Pagination(queryParams)).
			SelectAndCount()
		if err != nil {
			log.Println(err)
			return 0, nil, err
		}

		return count, vehicleList, nil
	}
	carType = strings.ToLower(carType)

	count, err := db.Model(&vehicleList).
		Apply(orm.Pagination(queryParams)).
		Where("vehicle_type = ?", carType).
		SelectAndCount()
	if err != nil {
		log.Println(err)
		return 0, nil, err
	}

	return count, vehicleList, nil

}

func GetCarsByType(queryParams url.Values, paramType string) (int, []models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	count, err := db.Model(&vehicleList).
		Where("vehicle_type = ?", paramType).
		Apply(orm.Pagination(queryParams)).
		SelectAndCount()
	if err != nil {
		log.Println(err)
		return 0, nil, err
	}

	return count, vehicleList, nil
}

func GetCarDetail(v models.Vehicle) (models.Vehicle, error) {
	db := connectToDB()

	if err := db.Model(&v).
		WherePK().
		Select(); err != nil {
		return models.Vehicle{}, err
	}

	return v, nil
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
		return []models.Vehicle{}, nil
	}

	return users[0].Vehicles, nil
}

func GetExplore() (map[string]interface{}, error) {
	vehicleMap := make(map[string]interface{})

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
		exploreMap := make(map[string]interface{})

		list, err = getTypeVehicleArray(carType)
		if err != nil {
			return nil, err
		}

		exploreMap["vehicles"] = list

		if carType == "suv" {
			displayName := strings.ToUpper(carType) + "'s"
			exploreMap["display_name"] = displayName
		}
		if carType == "sports" {
			displayName := carType + " cars"
			displayName = strings.Title(displayName)
			exploreMap["display_name"] = displayName
		}
		if carType == "sedan" {
			exploreMap["display_name"] = strings.Title(carType) + "s"
		}

		vehicleMap[carType] = exploreMap
	}

	return vehicleMap, nil
}

func getTypeVehicleArray(carType string) ([]models.Vehicle, error) {
	var list []models.Vehicle

	db := connectToDB()

	if err := db.
		Model(&list).
		Column("id", "make", "model", "year", "thumbnails", "vehicle_type", "daily_price").
		Where("vehicle_type = ?", carType).
		Limit(3).
		Select(); err != nil {
		return nil, err
	}

	return list, nil
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
		log.Println(err)
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
