package handlers

import (
	"net/http"
	"strconv"
	"time"

	"sirh/database"
	"sirh/models"

	"github.com/gin-gonic/gin"
)

func GetSalaires(c *gin.Context) {
	var salaires []models.SalaireEmploye
	if err := database.DB.Preload("Utilisateur").Order("date_debut desc").Find(&salaires).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salaires)
}

func CreateSalaire(c *gin.Context) {
	var input struct {
		UtilisateurID uint      `json:"utilisateur_id" binding:"required"`
		SalaireBase   float64   `json:"salaire_base" binding:"required"`
		DateDebut     time.Time `json:"date_debut" binding:"required"`
		DateFin       *time.Time `json:"date_fin"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.SalaireEmploye{
		UtilisateurID: input.UtilisateurID,
		SalaireBase:   input.SalaireBase,
		DateDebut:     input.DateDebut,
		DateFin:       input.DateFin,
	}

	if err := database.DB.Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Utilisateur").First(&s, s.ID)
	c.JSON(http.StatusCreated, s)
}

func UpdateSalaire(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var s models.SalaireEmploye
	if err := database.DB.First(&s, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Salaire non trouvé"})
		return
	}

	var input struct {
		SalaireBase *float64    `json:"salaire_base"`
		DateDebut   *time.Time  `json:"date_debut"`
		DateFin     **time.Time `json:"date_fin"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.SalaireBase != nil {
		s.SalaireBase = *input.SalaireBase
	}
	if input.DateDebut != nil {
		s.DateDebut = *input.DateDebut
	}
	if input.DateFin != nil {
		s.DateFin = *input.DateFin
	}

	if err := database.DB.Save(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Utilisateur").First(&s, s.ID)
	c.JSON(http.StatusOK, s)
}

func DeleteSalaire(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.SalaireEmploye{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Salaire supprimé"})
}

