package controllers

import (
	"net/http"
	"restaurante-api/database"
	"restaurante-api/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("MenuItems").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}
	if err := database.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la orden"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orden eliminada"})
}

func CreateTakeAwayOrder(c *gin.Context) {
	var input struct {
		PickupTime string `json:"pickup_time"`
		MenuIDs    []uint `json:"menu_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var menus []models.Menu
	if err := database.DB.Where("id IN ?", input.MenuIDs).Find(&menus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menús no encontrados"})
		return
	}

	var total float64
	var menuPtrs []*models.Menu
	for i := range menus {
		total += menus[i].Price
		menuPtrs = append(menuPtrs, &menus[i])
	}

	order := models.Order{
		Type:      "TakeAway",
		Total:     total,
		MenuItems: menuPtrs,
	}
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	takeAway := models.TakeAwayOrder{
		OrderID:    order.ID,
		PickupTime: input.PickupTime,
	}
	if err := database.DB.Create(&takeAway).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"order":    order,
		"takeaway": takeAway,
	})
}

func CreateShippingOrder(c *gin.Context) {
	var input struct {
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
		Status      string `json:"status"`
		MenuIDs     []uint `json:"menu_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var menus []models.Menu
	if err := database.DB.Where("id IN ?", input.MenuIDs).Find(&menus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menús no encontrados"})
		return
	}

	var total float64
	var menuPtrs []*models.Menu
	for i := range menus {
		total += menus[i].Price
		menuPtrs = append(menuPtrs, &menus[i])
	}

	order := models.Order{
		Type:      "Shipping",
		Total:     total,
		MenuItems: menuPtrs,
	}
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	shipping := models.ShippingOrder{
		OrderID:     order.ID,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Status:      input.Status,
	}
	if err := database.DB.Create(&shipping).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"order":    order,
		"shipping": shipping,
	})
}

func CreateEatInOrder(c *gin.Context) {
	var input struct {
		TableID uint   `json:"table_id"`
		MenuIDs []uint `json:"menu_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var menus []models.Menu
	if err := database.DB.Where("id IN ?", input.MenuIDs).Find(&menus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudieron obtener los menús"})
		return
	}

	var total float64
	var menuPtrs []*models.Menu
	for i := range menus {
		total += menus[i].Price
		menuPtrs = append(menuPtrs, &menus[i])
	}

	order := models.Order{
		Type:      "EatIn",
		Total:     total,
		MenuItems: menuPtrs,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	eatin := models.EatInOrder{
		OrderID: order.ID,
		TableID: input.TableID,
	}

	if err := database.DB.Create(&eatin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"order": order,
		"eatin": eatin,
	})
}
