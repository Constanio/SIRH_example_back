package services

import (
	"sirh/database"
	"sirh/models"
)

// DÉPARTEMENTS
func GetAllDepartements() ([]models.Departement, error) {
	var deps []models.Departement
	err := database.DB.Preload("Manager").Find(&deps).Error
	return deps, err
}

func CreateDepartement(d *models.Departement) error {
	return database.DB.Create(d).Error
}

func UpdateDepartement(d *models.Departement) error {
	return database.DB.Model(&models.Departement{}).Where("id = ?", d.ID).Updates(map[string]interface{}{
		"nom":         d.Nom,
		"code":        d.Code,
		"manager_id":  d.ManagerID,
		"description": d.Description,
	}).Error
}

func DeleteDepartement(id uint) error {
	return database.DB.Delete(&models.Departement{}, id).Error
}

// POSTES
func GetAllPostes() ([]models.Poste, error) {
	var postes []models.Poste
	err := database.DB.Preload("Departement").Find(&postes).Error
	return postes, err
}

func CreatePoste(p *models.Poste) error {
	return database.DB.Create(p).Error
}

func UpdatePoste(p *models.Poste) error {
	return database.DB.Model(&models.Poste{}).Where("id = ?", p.ID).Updates(map[string]interface{}{
		"titre":          p.Titre,
		"departement_id": p.DepartementID,
		"description":    p.Description,
		"salaire_min":    p.SalaireMin,
		"salaire_max":    p.SalaireMax,
	}).Error
}

func DeletePoste(id uint) error {
	return database.DB.Delete(&models.Poste{}, id).Error
}
