package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type OmdbAudit struct {
	ID             int32          `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	ClientPath     string         `json:"client_path"`
	ClientResponse postgres.Jsonb `json:"client_response"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      *time.Time     `json:"-"`
}

func (OmdbAudit) TableName() string {
	return "omdb_audit"
}
