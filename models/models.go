package models

import (
	"time"
)

type User struct { // postgres struct
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Record    uint
}

type Meta struct { // redis struct
	condition    string
	currentScore string
	answer       string
}
