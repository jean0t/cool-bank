package database

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role string `gorm:"not null"`
}

func (u *User) BeforeCreate(record *gorm.DB) error {
	if u.Role != "user" && u.Role != "manager" {
		return fmt.Errorf("invalid role: must be 'user' or 'manager'")
	}

	return nil
}

type Account struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	Owner User `gorm:"foreignKey:UserID"`
	AccountCode string `gorm:"uniqueIndex;not null;size:6"`
	ValueOwned float64 `gorm:"not null"`
}


type Operation struct {
	gorm.Model
	Value float64 `gorm:"not null"`
	FromID uint `gorm:"not null"`
	From Account `gorm:"foreignKey:FromID"`
	ToID uint `gorm:"not null"`
	To Account `gorm:"foreignKey:ToID"`
	Operation string `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
	Reverted bool
}
