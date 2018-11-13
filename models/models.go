package models

import "time"

type Vehicle struct {
	Address      string        `json:"address"`
	City         string        `json:"city"`
	Color        string        `json:"color"`
	Coordinates  []float64     `json:"coordinates" sql:",array"`
	CreatedTime  time.Time     `json:"created_time"`
	DayMax       int           `json:"day_max"`
	DayMin       int           `json:"day_min"`
	Description  string        `json:"description"`
	Doors        int           `json:"doors"`
	ID           string        `json:"id"`
	IsPublished  bool          `json:"is_published"`
	LicensePlate string        `json:"license_plate"`
	Make         string        `json:"make"`
	Miles        int           `json:"miles"`
	Model        string        `json:"model"`
	Photos       []string      `json:"photos" sql:",array"`
	Place        string        `json:"place"`
	Premium      bool          `json:"premium"`
	Seats        int           `json:"seats"`
	State        string        `json:"state"`
	Status       string        `json:"status" sql:"type:status"`
	Transmission string        `json:"transmission" sql:"type:transmission"`
	Trim         string        `json:"trim"`
	UpdatedBy    string        `json:"updated_by"`
	UpdatedTime  time.Time     `json:"updated_time"`
	UserID       string        `json:"user" sql:",fk"`
	VehicleNotes []VehicleNote `json:"notes" sql:",fk"`
	VehicleType  string        `json:"vehicle_type"`
	ViewIndex    int           `json:"view_index"`
	Vin          string        `json:"vin"`
	Year         int           `json:"year"`
	ZipCode      string        `json:"zip_code"`
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
	if target.Place != "" {
		source.Place = target.Place
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
	if target.ViewIndex != -1 {
		source.ViewIndex = target.ViewIndex
	}
	if target.Vin != "" {
		source.Vin = target.Vin
	}
	if target.Year != 0 {
		source.Year = target.Year
	}
	if target.UserID != "" {
		source.UserID = target.UserID
	}
	if target.ZipCode != "" {
		source.ZipCode = target.ZipCode
	}

	return source
}

type VehicleNote struct {
	Comment     string    `json:"comment"`
	CreatedBy   string    `json:"created_by"`
	CreatedTime time.Time `json:"created_time"`
	ID          int       `json:"id" sql:",pk"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedTime time.Time `json:"updated_time"`
	VehicleID   string    `json:"vehicle_id" sql:",fk"`
}

type UserNote struct {
	Comment     string    `json:"comment"`
	CreatedBy   string    `json:"created_by"`
	CreatedTime time.Time `json:"created_time"`
	ID          int       `json:"id" sql:",pk"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedTime time.Time `json:"updated_time"`
	UserID      string    `json:"user_id" sql:",fk"`
}

type User struct {
	ID        string     `json:"id" sql:",unique"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Vehicles  []Vehicle  `json:"vehicles" sql:",fk"`
	UserNotes []UserNote `json:"notes" sql:",fk"`
}

func (target *User) Merge(source User) User {
	if target.Email != "" {
		source.Email = target.Email
	}
	if target.Phone != "" {
		source.Phone = target.Phone
	}

	return source
}
