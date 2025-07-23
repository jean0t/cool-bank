package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)


func MigrateDB(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&Account{}, &Operations{})
	if err != nil {
		return err
	}

	return nil
}


func ConnectDB(dbName string) (*gorm.DB, error) {
	
}
