package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"

	h "github.com/omdbRestAPI/app/helper"
	"github.com/omdbRestAPI/config"
)

var initializeDB = false
var (
	db  *gorm.DB
	err error
)

func InitDB() {
	if !initializeDB {
		dbConnection := fmt.Sprintf("host=%[1]s port=%[3]s user=%[4]s dbname=%[2]s password=%[5]s", os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
		db, err = gorm.Open("postgres", dbConnection)

		if err != nil {
			h.Log(h.LogTypeFatal, "Failed connect to Database")
		}

		db.DB().SetMaxIdleConns(config.DatabaseMaxIdleConns)
		db.DB().SetMaxOpenConns(config.DatabaseMaxOpenConns)

		initializeDB = true
	}
}

func GetDB() *gorm.DB {
	return db
}
