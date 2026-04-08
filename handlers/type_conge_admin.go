package handlers

import (
	"net/http"
	"strconv"

	"sirh/database"
	"sirh/models"

	"github.com/gin-gonic/gin"
)

func CreateTypeConge(c *gin.Context) {
	var input models.TypeConge
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func UpdateTypeConge(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var tc models.TypeConge
	if err := database.DB.First(&tc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type de congé non trouvé"})
		return
	}

	var input struct {
		Nom                  *string `json:"nom"`
		Description          *string `json:"description"`
		JoursParAn           *int    `json:"jours_par_an"`
		NecessiteApprobation *bool   `json:"necessite_approbation"`
		Couleur              *string `json:"couleur"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nom != nil {
		tc.Nom = *input.Nom
	}
	if input.Description != nil {
		tc.Description = *input.Description
	}
	if input.JoursParAn != nil {
		tc.JoursParAn = *input.JoursParAn
	}
	if input.NecessiteApprobation != nil {
		tc.NecessiteApprobation = *input.NecessiteApprobation
	}
	if input.Couleur != nil {
		tc.Couleur = *input.Couleur
	}

	if err := database.DB.Save(&tc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tc)
}

func DeleteTypeConge(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Delete(&models.TypeConge{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Type de congé supprimé"})
}

