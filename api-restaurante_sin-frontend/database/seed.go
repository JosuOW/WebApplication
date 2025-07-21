package database

import (
	"log"
	"restaurante-api/models"

	"gorm.io/gorm"
)

func SeedMenus(db *gorm.DB) {
	menus := []models.Menu{
		{Name: "Pizza Margarita", Price: 9.99, Content: "Queso mozzarella, tomate fresco, albahaca", Active: true, Waiter: false},
		{Name: "Encebollado", Price: 4.75, Content: "Pescado, yuca, cebolla curtida, limón", Active: true, Waiter: false},
		{Name: "Jugo Natural de Maracuyá", Price: 2.00, Content: "Maracuyá, agua o leche, sin azúcar añadida", Active: true, Waiter: false},
		{Name: "Ceviche Mixto", Price: 6.50, Content: "Camarón, pescado, limón, cebolla, chifle", Active: true, Waiter: true},
		{Name: "Hamburguesa Criolla", Price: 7.20, Content: "Carne de res, huevo frito, plátano maduro", Active: true, Waiter: false},
	}

	for _, menu := range menus {
		var existing models.Menu
		if err := db.Where("name = ?", menu.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&menu).Error; err != nil {
				log.Printf("❌ Error creando menú %s: %v", menu.Name, err)
			} else {
				log.Printf("✅ Menú creado: %s", menu.Name)
			}
		} else {
			log.Printf("⚠️  Menú ya existe: %s", menu.Name)
		}
	}
}

func SeedTables(db *gorm.DB) {
	tables := []models.Table{
		{Number: 1, Seats: 4, Occupied: false, Location: "Interior"},
		{Number: 2, Seats: 2, Occupied: false, Location: "Exterior"},
		{Number: 3, Seats: 6, Occupied: false, Location: "Interior"},
		{Number: 4, Seats: 2, Occupied: false, Location: "Terraza"},
		{Number: 5, Seats: 4, Occupied: false, Location: "Exterior"},
		{Number: 6, Seats: 8, Occupied: false, Location: "VIP"},
	}

	for _, table := range tables {
		var existing models.Table
		if err := db.Where("number = ?", table.Number).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&table).Error; err != nil {
				log.Printf("❌ Error creando mesa %d: %v", table.Number, err)
			} else {
				log.Printf("✅ Mesa creada: %d", table.Number)
			}
		}
	}
}
