package handlers

import (
	"sirh/models"
	"sirh/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DÉPARTEMENTS
func GetDepartements(c *gin.Context) {
	deps, err := services.GetAllDepartements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deps)
}

func CreateDepartement(c *gin.Context) {
	var input models.Departement
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateDepartement(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func UpdateDepartement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Departement
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = uint(id)
	if err := services.UpdateDepartement(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, input)
}

func DeleteDepartement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteDepartement(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Département supprimé"})
}

// POSTES
func GetPostes(c *gin.Context) {
	postes, err := services.GetAllPostes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, postes)
}

func CreatePoste(c *gin.Context) {
	var input models.Poste
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreatePoste(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func UpdatePoste(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Poste
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = uint(id)
	if err := services.UpdatePoste(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, input)
}

func DeletePoste(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeletePoste(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Poste supprimé"})
}
