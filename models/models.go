package models

import "time"

type Vehicle struct {
	ID           string    `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	Trim         string    `json:"trim"`
	Color        string    `json:"color"`
	Doors        int       `json:"doors"`
	Seats        int       `json:"seats"`
	Vin          string    `json:"vin"`
	Description  string    `json:"description"`
	DayMax       int       `json:"day_max"`
	DayMin       int       `json:"day_min"`
	VehicleType  string    `json:"vehicle_type"`
	Photos       []string  `json:"photos" sql:",array"`
	Miles        int       `json:"miles"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedTime  time.Time `json:"updated_time"`
	User         string    `json:"user_sub"`
	IsPublished  bool      `json:"is_published"`
	IsApproved   bool      `json:"is_approved"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Coordinates  []float64 `json:"coordinates" sql:",array"`
}

type VehicleSignupStage struct {
	Stage     int    `json:"stage"`
	User      string `json:"user_sub"`
	VehicleID string `json:"vehicle_id" pg:",fk:vehicle_id"`
	Completed bool   `json:"completed"`
}

type UserInfo struct {
	tableName    struct{} `sql:"user_info"`
	Sub          string   `json:"sub" sql:",pk,unique"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	ListedCars   []string `json:"listed_cars" sql:",array"`
	UnlistedCars []string `json:"unlisted_cars" sql:",array"`
}
