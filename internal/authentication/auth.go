package authentication

import (
	"fmt"
	"time"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

//==========================================| Types

type contextKey string

type Claims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}


//==========================================| Funtions

func CreateCookie(name, role string, expiration time.Duration, pathPrivateKey string) (string, error) {
	var (
		err error
		claims *Claims
		tokenString string
		token *jwt.Token
		privateKey *rsa.PrivateKey
		
		now time.Time
	)

	now = time.Now()
	claims = &Claims{
			Name: name,
			Role: role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(now.Add(expiration)),
				IssuedAt: jwt.NewNumericDate(now),
				},
			}
	
	privateKey, err = loadPrivateKey(pathPrivateKey)
	if err != nil {
		return "", err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err = token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func VerifyCookie(tokenString, pathPublicKey string) bool {
	var (
		err error
		ok, isValid bool
		publicKey *rsa.PublicKey
		token *jwt.Token
	)

	publicKey, err = loadPublicKey(pathPublicKey)
	if err != nil {
		fmt.Println("[!] Error loading RSA Public Key")
		return false
	}

	token, err = jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if _, ok = token.Claims.(*Claims); ok && token.Valid {
		isValid = true
	} else {
		isValid =  false
	}

	return isValid
}
