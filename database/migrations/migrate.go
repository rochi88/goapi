package migrations

import (
	"log"

	"github.com/rochi88/goapi/app/helpers"
	"github.com/rochi88/goapi/app/models"
)

func Migrate() {
	if err := helpers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migrated!")
}
