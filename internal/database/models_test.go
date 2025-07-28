package database

import (
	"gorm.io/gorm"
	
	"testing"
	"context"
)


func TestUserCreation(t *testing.T) {
	var (
		db *gorm.DB
		err error
		ctx context.Context
	)
	db = setupTestDB(t)

	ctx = context.Background()
	var tests = []struct{
		title string
		username string
		password string
		role string
		expectError bool
	}{
		{
			title: "valid user",
			username: "user test",
			password: "any_secret",
			role: "user",
			expectError: false,
		},
		{
			title: "valid manager",
			username: "manager test",
			password: "another_secret",
			role: "manager",
			expectError: false,
		},
		{
			title: "invalid user",
			username: "failed test",
			password: "anyway",
			role: "unknown",
			expectError: true,
		},
		{
			title: "mistyped role",
			username: "failed test two",
			password: "anyway_error",
			role: "Manager",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			var user = User{
				Username: tt.username,
				Password: tt.password,
				Role: tt.role,
			}
			err = gorm.G[User](db).Create(ctx, &user)
			if tt.expectError {
				if err == nil {
					t.Fatal("expected error, but didn't fail")
				}
			}
		})
	}
}
