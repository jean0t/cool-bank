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


type Operation struct {
	gorm.Model
	Value float64 `gorm:"not null"`
	FromID uint `gorm:"not null"`
	From Account `gorm:"foreignKey:FromID"`
	ToID uint `gorm:"not null"`
	To Account `gorm:"foreignKey:ToID"`
	Operation string `gorm:"not null"`
}
