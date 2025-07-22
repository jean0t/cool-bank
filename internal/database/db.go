package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/sqlite"
)

func MigrateDB(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&Account{}, &Operations{})
	if err != nil {
		return err
	}

	return nil
}
