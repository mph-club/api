package server

import "time"

type Vehicle struct {
	ID                  string    `json:"id"`
	Make                string    `json:"make"`
	Model               string    `json:"model"`
	Year                int       `json:"year"`
	Trim                string    `json:"trim"`
	Color               string    `json:"color"`
	Doors               int       `json:"doors"`
	Seats               int       `json:"seats"`
	Vin                 string    `json:"vin"`
	Description         string    `json:"description"`
	DayMax              int       `json:"day_max"`
	DayMin              int       `json:"day_min"`
	VehicleType         string    `json:"vehicle_type"`
	Photos              []string  `json:"photos"`
	VehicleRegistration string    `json:"vehicle_registration"`
	Status              string    `json:"status"`
	CreatedBy           string    `json:"created_by"`
	CreatedTime         time.Time `json:"created_time"`
	UpdatedBy           string    `json:"updated_by"`
	UpdatedTime         time.Time `json:"updated_time"`
}
