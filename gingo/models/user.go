package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Bio  string
	Age  uint8
}
