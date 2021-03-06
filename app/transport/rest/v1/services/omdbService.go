package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/omdbRestAPI/app/helper"
	"github.com/omdbRestAPI/app/transport/rest/v1/models"
	"github.com/omdbRestAPI/app/transport/rest/v1/repository"
	"time"
)

func GetMoviesByParam(path string) (map[string]interface{}, error) {
	var (
		err error
	)
	responseMap, err := helper.NewRequest("GET", path, nil, nil, helper.BasicAuth{})

	if err != nil {
		return nil, err
	}

	clientResponse, _ := json.Marshal(responseMap)

	auditData := models.OmdbAudit{
		ClientPath:     path,
		ClientResponse: postgres.Jsonb{clientResponse},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
	}

	go repository.CreateOmdbAudit(&auditData)

	return responseMap, nil
}
