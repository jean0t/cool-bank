package database

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Phone string
	ValueOwned float64
}


type Operations struct {
	gorm.Model
	Value string `gorm:"not null"`
	From Account `gorm:"not null"`
	To Account `gorm:"not null"`
	Operation string `gorm:"not null"`
}
