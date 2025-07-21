package authentication

import (
	"io"
	"os"
	"fmt"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
)

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	var (
		err error
		file *os.File
		keyData []byte
		privateKey *rsa.PrivateKey
	)

	file, err = os.Open(path)
	if err != nil {
		fmt.Println("[!] Error opening RSA Private Key.")
		return nil, err
	}
	defer file.Close()

	keyData, err = io.ReadAll(file)
	if err != nil {
		fmt.Println("[!] Error loading RSA Private Key.")
		return nil, err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		fmt.Println("[!] Error Parsing RSA Private Key.")
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	var (
		err error
		file *os.File
		keyData []byte
		publicKey *rsa.PublicKey
	)

	file, err = os.Open(path)
	if err != nil {
		fmt.Println("[!] Error opening RSA Public Key.")
		return nil, err
	}
	defer file.Close()

	keyData, err = io.ReadAll(file)
	if err != nil {
		fmt.Println("[!] Error reading RSA Public Key.")
		return nil, err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromREM(keyData)
	if err != nil {
		fmt.Println("[!] Error parsing RSA Public Key.")
		return nil, err
	}

	return publicKey, nil
}
