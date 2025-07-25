package utils

import (
	"math/rand"
	"time"

	"github.com/jean0t/cool-bank/internal/database"

	"gorm.io/gorm"
)

const accountNumberLength = 6
const charset string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomCode(length int) string {
	var (
		code []byte
	)
	
	code = make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(code)
}


func GenerateUniqueAccountNumber(db *gorm.DB, maxRetries int) (string, error) {
	var (
		accountCode string
		count int64
		err error
	)

	for i := 0; i < maxRetries; i++ {
		err = db.Model(&database.Account{}).Where("AccountCode = ?", accountCode).Count(&count).Error
		if err != nil {
			return "", err
		}

		if count == 0 {
			return accountCode, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique account code")
}
