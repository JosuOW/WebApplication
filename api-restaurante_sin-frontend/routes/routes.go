package routes

import (
	"restaurante-api/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/menus", controllers.GetMenus)
	r.POST("/menus", controllers.CreateMenu)
	r.DELETE("/menus/:id", controllers.DeleteMenu)

	r.GET("/orders", controllers.GetOrders)
	r.POST("/orders/takeaway", controllers.CreateTakeAwayOrder)
	r.POST("/orders/shipping", controllers.CreateShippingOrder)
	r.POST("/orders/eatin", controllers.CreateEatInOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	return r
}
