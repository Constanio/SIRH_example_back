package handlers

import (
	"net/http"
	"strconv"
	"time"

	"sirh/database"
	"sirh/models"

	"github.com/gin-gonic/gin"
)

func GetFichesPaie(c *gin.Context) {
	var fiches []models.FichePaie
	if err := database.DB.Preload("Utilisateur").Order("annee desc, mois desc").Find(&fiches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fiches)
}

func GetFichePaie(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var fp models.FichePaie
	if err := database.DB.Preload("Utilisateur").First(&fp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fiche de paie non trouvée"})
		return
	}
	c.JSON(http.StatusOK, fp)
}

func CreateFichePaie(c *gin.Context) {
	var input struct {
		UtilisateurID uint            `json:"utilisateur_id" binding:"required"`
		Mois          int             `json:"mois" binding:"required"`
		Annee         int             `json:"annee" binding:"required"`
		SalaireBase   float64         `json:"salaire_base"`
		Primes        float64         `json:"primes"`
		Deductions    float64         `json:"deductions"`
		DatePaiement  *time.Time      `json:"date_paiement"`
		Statut        models.StatutPaie `json:"statut"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fp := models.FichePaie{
		UtilisateurID: input.UtilisateurID,
		Mois:          input.Mois,
		Annee:         input.Annee,
		SalaireBase:   input.SalaireBase,
		Primes:        input.Primes,
		Deductions:    input.Deductions,
		DatePaiement:  input.DatePaiement,
	}
	if input.Statut != "" {
		fp.Statut = input.Statut
	}

	if err := database.DB.Create(&fp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Utilisateur").First(&fp, fp.ID)
	c.JSON(http.StatusCreated, fp)
}

func UpdateFichePaie(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var fp models.FichePaie
	if err := database.DB.First(&fp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fiche de paie non trouvée"})
		return
	}

	var input struct {
		SalaireBase  *float64         `json:"salaire_base"`
		Primes       *float64         `json:"primes"`
		Deductions   *float64         `json:"deductions"`
		DatePaiement **time.Time      `json:"date_paiement"`
		Statut       *models.StatutPaie `json:"statut"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.SalaireBase != nil {
		fp.SalaireBase = *input.SalaireBase
	}
	if input.Primes != nil {
		fp.Primes = *input.Primes
	}
	if input.Deductions != nil {
		fp.Deductions = *input.Deductions
	}
	if input.DatePaiement != nil {
		fp.DatePaiement = *input.DatePaiement
	}
	if input.Statut != nil {
		fp.Statut = *input.Statut
	}

	if err := database.DB.Save(&fp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Utilisateur").First(&fp, fp.ID)
	c.JSON(http.StatusOK, fp)
}

func DeleteFichePaie(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.FichePaie{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fiche de paie supprimée"})
}

