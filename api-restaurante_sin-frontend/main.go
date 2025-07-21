package main

import (
	"restaurante-api/database"
	"restaurante-api/models"
	"restaurante-api/routes"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(
		&models.Menu{},
		&models.Customer{},
		&models.Table{},
		&models.Booking{},
		&models.Order{},
		&models.TakeAwayOrder{},
		&models.ShippingOrder{},
		&models.EatInOrder{},
		&models.OrderMenu{},
	)

	database.SeedMenus(database.DB)
	database.SeedTables(database.DB)

	r := routes.SetupRoutes()
	r.Run(":8080")
}
