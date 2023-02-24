package storage

import (
	config "ProGolang/configs"
	models "ProGolang/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func PostgreSQL() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetPostgresHost(), config.PostgresPort(), config.GetPostgresUser(), config.GetPostgresPassword(), config.PostgresDataBaseName(), config.PostgresSSLMode(),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate() {
	log.Printf("Migrate: Start")

	db := PostgreSQL()
	err := models.MigrateUser(db)
	if err != nil {
		panic("failed migrate user table")
	}

	log.Printf("Migrate: Success")
}
