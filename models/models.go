package models

import (
	"gorm.io/gorm"
)

type User struct { // postgres struct
	gorm.Model
	Name   string
	Record string
}

type Meta struct { // redis struct
	condition    string
	currentScore string
	answer       string
}
