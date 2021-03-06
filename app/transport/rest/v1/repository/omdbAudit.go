package repository

import "github.com/omdbRestAPI/app/transport/rest/v1/models"

func CreateOmdbAudit(omdbAudit *models.OmdbAudit) error {
	db := getDB()

	if err := db.Create(omdbAudit).Error; err != nil {
		return err
	}
	return nil
}
