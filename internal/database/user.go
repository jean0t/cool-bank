package database

import (
	"fmt"
	"time"
	"context"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(db *gorm.DB, username, password, role string) error {
	var (
		user User
		hashedPassword []byte
		hashed string
		err error
		ctx context.Context
		cancel context.CancelFunc
	)
	
	// if query takes longer than 2s an error will be produced
	ctx, cancel = context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel() // if query takes less than 2s the resources will be released

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// create the user
	hashed = fmt.Sprintf("%x", hashedPassword)
	user = User{Username: username, Password: hashed, Role, role}
	err = gorm.G[User](db).Create(ctx, &user)

	return err
}


func VerifyUser(db *gorm.DB, username, password string) error {
	var (
		ctx context.Context
		err error
		user User
	)
	ctx = context.Background()
	user, err = gorm.G[User](db).Where("username = ?", username).First(ctx)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}
