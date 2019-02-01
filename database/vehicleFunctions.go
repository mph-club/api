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

func EditPhotoURLArrayOnVehicle(vehicleID, filename string) error {
	db = connectToDB()

	url := fmt.Sprintf("https://mphclub.s3.amazonaws.com/%s/%s", vehicleID, filename)
	thumbnailURL := fmt.Sprintf("https://mphclub.s3.amazonaws.com/%s/thumb/%s", vehicleID, filename)

	vehicleToAttach := &models.Vehicle{
		ID: vehicleID,
	}

	err := db.Select(vehicleToAttach)
	if err != nil {
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
		return err
	}

	return nil
}

func UpsertListing(v models.Vehicle) (string, string, error) {
	db = connectToDB()

	car := models.Vehicle{
		ID: v.ID,
	}

	if err := db.Select(&car); err != nil {
		log.Println(err.Error())
		log.Println("car does not exist, create")
	} else {
		log.Println("car does exist, update")

		feature := v.Feature

		v = v.Merge(car)
		v.UpdatedTime = time.Now()

		if dbErr := db.Update(&v); dbErr != nil {
			return "", "", dbErr
		}

		if feature != car.Feature {
			if err := insertFeatureAndUpdateVehicle(feature, v.ID); err != nil {
				return "", "", err
			}
		}

		return v.ID, "updated", nil
	}

	v.ID = xid.New().String()

	if v.ViewIndex == 0 {
		v.ViewIndex = -1
	}

	v.CreatedTime = time.Now()
	v.Status = "PENDING"

	if err := db.Insert(&v); err != nil {
		return "", "", err
	}

	if err := insertFeatureAndUpdateVehicle(v.Feature, v.ID); err != nil {
		return "", "", err
	}

	log.Println("vehicle created")
	return v.ID, "created", nil
}

func insertFeatureAndUpdateVehicle(feature models.Features, vehicleID string) error {
	feature.VehicleID = vehicleID

	if err := db.Insert(&feature); err != nil {
		return err
	}

	var v models.Vehicle
	v.ID = vehicleID
	v.FeatureID = feature.ID

	_, err := db.Model(&v).
		Column("feature_id").
		Where("id = ?", v.ID).
		Update()
	if err != nil {
		return err
	}

	return nil
}

func GetCars(queryParams url.Values, carType string) (int, []models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db = connectToDB()

	if len(carType) == 0 {
		count, err := db.Model(&vehicleList).
			Apply(orm.Pagination(queryParams)).
			Where("status = ?", "APPROVED").
			SelectAndCount()
		if err != nil {
			return 0, nil, err
		}

		return count, vehicleList, nil
	}
	carType = strings.ToLower(carType)

	count, err := db.Model(&vehicleList).
		Apply(orm.Pagination(queryParams)).
		Where("vehicle_type = ?", carType).
		Where("status = ?", "APPROVED").
		SelectAndCount()
	if err != nil {
		return 0, nil, err
	}

	return count, vehicleList, nil

}

func GetCarsByType(queryParams url.Values, paramType string) (int, []models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db = connectToDB()

	count, err := db.Model(&vehicleList).
		Where("vehicle_type = ?", paramType).
		Apply(orm.Pagination(queryParams)).
		SelectAndCount()
	if err != nil {
		return 0, nil, err
	}

	return count, vehicleList, nil
}

func GetCarDetail(v models.Vehicle) (models.Vehicle, error) {
	db = connectToDB()

	var vArray []models.Vehicle

	if err := db.Model(&vArray).
		Where("vehicle.id = ?", v.ID).
		Column("vehicle.*", "Feature").
		Relation("Feature").
		Select(); err != nil {
		return models.Vehicle{}, err
	}

	v = vArray[0]

	owner, err := GetUser(v.UserID)
	if err != nil {
		return v, err
	}

	v.OwnerDetails = map[string]interface{}{
		"first_name":        owner.DriverLicense.FirstName,
		"rating":            5,
		"response_rate":     90,
		"response_time":     5,
		"profile_photo_url": owner.ProfilePhotoURL,
	}

	return v, nil
}

func YouAlsoMightLike(vehicleType string) ([]models.Vehicle, error) {
	list, err := getTypeVehicleArray(vehicleType)
	if err != nil {
		return []models.Vehicle{}, err
	}

	return list, nil
}

func GetMyCars(u *models.User) ([]models.Vehicle, error) {
	db = connectToDB()

	var users []models.User

	err := db.Model(&users).
		Column("user.*", "Vehicles").
		Relation("Vehicles", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Where("id = ?", u.ID).
		Select()

	if err != nil {
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

	db = connectToDB()

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

		if carType == "sports" {
			displayName := carType + " cars"
			displayName = strings.Title(displayName)
			exploreMap["display_name"] = displayName
			exploreMap["order"] = 1
		}
		if carType == "sedan" {
			exploreMap["display_name"] = strings.Title(carType) + "s"
			exploreMap["order"] = 2
		}
		if carType == "suv" {
			displayName := strings.ToUpper(carType) + "'s"
			exploreMap["display_name"] = displayName
			exploreMap["order"] = 3
		}

		vehicleMap[carType] = exploreMap
	}

	return vehicleMap, nil
}

func getTypeVehicleArray(carType string) ([]models.Vehicle, error) {
	var list []models.Vehicle
	var listWithSingleThumb []models.Vehicle

	db = connectToDB()

	if err := db.
		Model(&list).
		Column("id", "make", "model", "year", "thumbnails", "vehicle_type", "daily_price").
		Where("vehicle_type = ?", carType).
		Where("status = ?", "APPROVED").
		Limit(3).
		Select(); err != nil {
		return nil, err
	}

	for _, vehicle := range list {
		vehicle.Thumbnails = []string{vehicle.Thumbnails[0]}
		listWithSingleThumb = append(listWithSingleThumb, vehicle)
	}

	return listWithSingleThumb, nil
}

func MakeReservation(trip models.Trip) error {
	db = connectToDB()

	if err := db.Insert(&trip); err != nil {
		return err
	}

	return nil
}

func GetUnavailableDates(vehicleID string) ([]interface{}, error) {
	db = connectToDB()

	var trips []models.Trip
	today := time.Now()

	unavailableDates := []interface{}{}

	if err := db.
		Model(&trips).
		Where("vehicle_ID = ?", vehicleID).
		Order("start_time DESC").
		Select(); err != nil {
		return unavailableDates, err
	}

	for _, trip := range trips {
		loc, _ := time.LoadLocation("America/New_York")

		start := trip.StartTime.In(loc)
		end := trip.EndTime.In(loc)

		if !end.Before(today) {
			unavailableObject := map[string]interface{}{
				"start": start,
				"end":   end,
			}

			unavailableDates = append(unavailableDates, unavailableObject)
		}
	}

	return unavailableDates, nil
}
