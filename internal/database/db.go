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
	var (
		db *gorm.DB
		driver gorm.Dialector
		err error
	)
	driver = sqlite.Open(dbName)
	db, err = gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
