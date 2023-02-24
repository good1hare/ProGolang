package modelas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `gorm:"primary key;autoIncrement" json:"id"`
	UserName string `json:"user_name"`
	ChatId   int64  `json:"chat_id"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
