package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Number   int    `json:"number"`
	Seats    int    `json:"seats"`
	Occupied bool   `json:"occupied"`
	Location string `json:"location"`

	Bookings []Booking
}
