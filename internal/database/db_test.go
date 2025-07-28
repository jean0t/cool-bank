package database

import (
	"gorm.io/gorm"
	"testing"
)

func TestConnectDB(t *testing.T) {
	var (
		err error
	)

	_, err = ConnectDB(":memory:")
	if err != nil {
		t.Fatal("Connection to db failed")
	}
}

func TestMigration(t *testing.T) {
	var (
		db *gorm.DB
		err error
	)

	db, _ = ConnectDB(":memory:")

	err = MigrateDB(db)
	if err != nil {
		t.Fatal("Failed to make migration")
	}
}
