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
	Status       string    `json:"status" sql:"type:status"`
	CreatedTime  time.Time `json:"created_time"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedTime  time.Time `json:"updated_time"`
	User         string    `json:"user_sub"`
	IsPublished  bool      `json:"is_published"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Coordinates  []float64 `json:"coordinates" sql:",array"`
	Notes        []string  `json:"notes" sql:",array"`
	ViewIndex    int       `json:"view_index"`
}

func (target *Vehicle) Merge(source Vehicle) Vehicle {
	if target.Address != "" {
		source.Address = target.Address
	}
	if target.City != "" {
		source.City = target.City
	}
	if target.Color != "" {
		source.Color = target.Color
	}
	if len(target.Coordinates) > 0 {
		source.Coordinates = target.Coordinates
	}
	if target.DayMax != 0 {
		source.DayMax = target.DayMax
	}
	if target.DayMin != 0 {
		source.DayMin = target.DayMin
	}
	if target.Description != "" {
		source.Description = target.Description
	}
	if target.Doors != 0 {
		source.Doors = target.Doors
	}
	if target.IsPublished {
		source.IsPublished = target.IsPublished
	}
	if target.LicensePlate != "" {
		source.LicensePlate = target.LicensePlate
	}
	if target.Make != "" {
		source.Make = target.Make
	}
	if target.Miles != 0 {
		source.Miles = target.Miles
	}
	if target.Model != "" {
		source.Model = target.Model
	}
	if len(target.Notes) > 0 {
		source.Notes = target.Notes
	}
	if target.Seats != 0 {
		source.Seats = target.Seats
	}
	if target.State != "" {
		source.State = target.State
	}
	if target.Status != "" {
		source.Status = target.Status
	}
	if target.Trim != "" {
		source.Trim = target.Trim
	}
	if target.VehicleType != "" {
		source.VehicleType = target.VehicleType
	}
	if target.ViewIndex != 0 {
		source.VehicleType = target.VehicleType
	}
	if target.Vin != "" {
		source.Vin = target.Vin
	}
	if target.Year != 0 {
		source.Year = target.Year
	}
	if target.User != "" {
		source.User = target.User
	}

	return source
}

type User struct {
	Sub          string   `json:"sub" sql:",pk,unique"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	ListedCars   []string `json:"listed_cars" sql:",array"`
	UnlistedCars []string `json:"unlisted_cars" sql:",array"`
}

func (target *User) Merge(source User) User {
	if target.Email != "" {
		source.Email = target.Email
	}
	if target.Phone != "" {
		source.Phone = target.Phone
	}
	if len(target.ListedCars) > 0 {
		source.ListedCars = append(source.ListedCars, target.ListedCars...)
	}
	if len(target.UnlistedCars) > 0 {
		source.UnlistedCars = append(source.UnlistedCars, target.UnlistedCars...)
	}

	return source
}
