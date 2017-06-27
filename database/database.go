package database

import (
	"os"

	"github.com/bgpat/tweet-picker/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	defaultDB *Database = nil
)

type Database struct {
	*gorm.DB
}

func New() (*Database, error) {
	url := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	models.AutoMigrate(db)
	res := &Database{
		DB: db,
	}
	if defaultDB == nil {
		defaultDB = res
	}
	return res, nil
}

func Default() *Database {
	if defaultDB == nil {
		New()
	}
	return defaultDB
}
