package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint
	UserName string
	ChatId   uint
}
