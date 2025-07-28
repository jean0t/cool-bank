package database

import (
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	var (
		db *gorm.DB
		err error
	)

	db, err = ConnectDB(":memory:")
	if err != nil {
		t.Fatal("Connection to db failed")
	}

	err = MigrateDB(db)
	if err != nil {
		t.Fatal("Failed to migrate the Database")
	}

	return db
}
