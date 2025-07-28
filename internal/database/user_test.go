package database

import (
	"gorm.io/gorm"

	"testing"
)


func queryUser(t *testing.T, db *gorm.DB, username string) int {
	t.Helper()
	var (
		err error
		count int64
	)
	err = db.Model(&User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		t.Fatal("Error counting the entry in the database")
	}

	return int(count)
}


func TestCreationOfUser(t *testing.T) {
	var (
		db *gorm.DB
		err error
		count int
	)

	db = setupTestDB(t)
	var tests = []struct {
		title string
		name string
		password string
		role string
		expectFail bool
	} {
		{
			title: "Valid user",
			name: "test valid",
			password: "strong_secret",
			role: "user",
			expectFail: false,
		},
		{
			title: "Invalid user",
			name: "test valid",
			password: "strong_secret",
			role: "failed",
			expectFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err = CreateUser(db, tt.name, tt.password, tt.role)
			if tt.expectFail && err == nil {
				t.Fatal("Expected error, but didn't fail")
			}

			count = queryUser(t, db, tt.name)
			if count < 1 && !tt.expectFail {
				t.Fatal("It failed to create User")
			}
		})

	}
}


func TestVerifyUser(t *testing.T) {
	var (
		err error
		db *gorm.DB
	)

	db = setupTestDB(t)

	var tests = []struct {
		title string
		name string
		password string
		passwordTest string
		role string
		expectFail bool
	}{
		{
			title: "valid password",
			name: "user_test1",
			password: "big_secret",
			passwordTest: "big_secret",
			role: "user",
			expectFail: false,

		},
		{
			title: "invalid password",
			name: "user_test2",
			password: "weak_secret",
			passwordTest: "strong_secret",
			role: "user",
			expectFail: true,

		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			_ = CreateUser(db, tt.name, tt.password, tt.role)
			err = VerifyUser(db, tt.name, tt.passwordTest)
			if tt.expectFail && err == nil {
				t.Fatal("Failed test, expected fail")
			}
		})
	}
}
