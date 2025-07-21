package main

import (
	"github.com/JosuOW/gestion-restaurante/api"
	"github.com/JosuOW/gestion-restaurante/database"
)

func main() {
	database.InitDB()
	router := api.SetupRouter()
	router.Run(":8080")
}
