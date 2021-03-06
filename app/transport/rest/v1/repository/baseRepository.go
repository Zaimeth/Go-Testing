package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/omdbRestAPI/app/core"
)

func getDB() *gorm.DB {
	return core.GetDB()
}
