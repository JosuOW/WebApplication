package models

import "gorm.io/gorm"

type OrderMenu struct {
	gorm.Model
	OrderID uint
	MenuID  uint
}
