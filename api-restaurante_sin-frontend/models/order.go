package models

import "gorm.io/gorm"

// Orden general
type Order struct {
	gorm.Model
	Type      string  `json:"type"`
	Total     float64 `json:"total"`
	MenuItems []*Menu `gorm:"many2many:order_menus"` // Relación con menús
}

// Orden para recoger en el restaurante
type TakeAwayOrder struct {
	gorm.Model
	OrderID    uint
	Order      Order
	PickupTime string `json:"pickup_time"`
}

// Orden para entrega a domicilio
type ShippingOrder struct {
	gorm.Model
	OrderID     uint
	Order       Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

// Orden para comer en el restaurante
type EatInOrder struct {
	gorm.Model
	OrderID uint  `json:"order_id"`
	Order   Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TableID uint  `json:"table_id"`
	Table   Table `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
