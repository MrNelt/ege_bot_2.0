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

type Token struct { // redis struct
	Condition    string `json:"condition"`
	CurrentScore uint   `json:"currentScore"`
	Record       uint   `json:"record"`
	Answer       string `json:"answer"`
}
