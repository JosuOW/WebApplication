package controllers

import (
	"net/http"
	"restaurante-api/database"
	"restaurante-api/models"

	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) {
	var menus []models.Menu
	database.DB.Find(&menus)
	c.JSON(http.StatusOK, menus)
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, menu)
}

func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	if err := database.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Elemento del menú no encontrado"})
		return
	}
	if err := database.DB.Delete(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el menú"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Elemento del menú eliminado correctamente"})
}
