package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	CustomerID uint
	TableID    uint
	DateTime   string `json:"date_time"`
	Guests     int    `json:"guests"`

	Customer Customer
	Table    Table
}
