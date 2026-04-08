package handlers

import (
	"net/http"
	"strconv"
	"time"

	"sirh/database"
	"sirh/models"

	"github.com/gin-gonic/gin"
)

func GetEvaluations(c *gin.Context) {
	var evals []models.EvaluationPerformance
	if err := database.DB.Preload("Utilisateur").Preload("Evaluateur").Order("date_evaluation desc").Find(&evals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, evals)
}

func CreateEvaluation(c *gin.Context) {
	var input struct {
		UtilisateurID  uint      `json:"utilisateur_id" binding:"required"`
		EvaluateurID   uint      `json:"evaluateur_id" binding:"required"`
		PeriodeDebut   time.Time `json:"periode_debut" binding:"required"`
		PeriodeFin     time.Time `json:"periode_fin" binding:"required"`
		DateEvaluation time.Time `json:"date_evaluation" binding:"required"`
		Score          float64   `json:"score"`
		Commentaires   string    `json:"commentaires"`
		Objectifs      string    `json:"objectifs"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ev := models.EvaluationPerformance{
		UtilisateurID:  input.UtilisateurID,
		EvaluateurID:   input.EvaluateurID,
		PeriodeDebut:   input.PeriodeDebut,
		PeriodeFin:     input.PeriodeFin,
		DateEvaluation: input.DateEvaluation,
		Score:          input.Score,
		Commentaires:   input.Commentaires,
		Objectifs:      input.Objectifs,
	}

	if err := database.DB.Create(&ev).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Utilisateur").Preload("Evaluateur").First(&ev, ev.ID)
	c.JSON(http.StatusCreated, ev)
}

func UpdateEvaluation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var ev models.EvaluationPerformance
	if err := database.DB.First(&ev, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Évaluation non trouvée"})
		return
	}

	var input struct {
		PeriodeDebut   *time.Time `json:"periode_debut"`
		PeriodeFin     *time.Time `json:"periode_fin"`
		DateEvaluation *time.Time `json:"date_evaluation"`
		Score          *float64   `json:"score"`
		Commentaires   *string    `json:"commentaires"`
		Objectifs      *string    `json:"objectifs"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.PeriodeDebut != nil {
		ev.PeriodeDebut = *input.PeriodeDebut
	}
	if input.PeriodeFin != nil {
		ev.PeriodeFin = *input.PeriodeFin
	}
	if input.DateEvaluation != nil {
		ev.DateEvaluation = *input.DateEvaluation
	}
	if input.Score != nil {
		ev.Score = *input.Score
	}
	if input.Commentaires != nil {
		ev.Commentaires = *input.Commentaires
	}
	if input.Objectifs != nil {
		ev.Objectifs = *input.Objectifs
	}

	if err := database.DB.Save(&ev).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Utilisateur").Preload("Evaluateur").First(&ev, ev.ID)
	c.JSON(http.StatusOK, ev)
}

func DeleteEvaluation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.EvaluationPerformance{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Évaluation supprimée"})
}

