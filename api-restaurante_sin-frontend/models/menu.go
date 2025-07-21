package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Content string  `json:"content"`
	Active  bool    `json:"active"`
	Waiter  bool    `json:"waiter"`

	Orders []*Order `gorm:"many2many:order_menus"`
}
