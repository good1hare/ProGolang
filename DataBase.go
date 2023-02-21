package main

import (
	config "ProGolang/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() {
	dsn := "host=localhost user=" + config.GetPostgresUser() + " password=" + config.GetPostgresPassword() + " dbname=" + config.PostgresDataBaseName() + " port=" + config.PostgresPort() + " sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrations(db)
}

type User struct {
	gorm.Model
	Id       uint
	UserName string
	ChatId   uint
}

// Миграция схем
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}

}
