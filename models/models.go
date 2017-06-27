package models

import (
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Tweet{}, &User{})
}
