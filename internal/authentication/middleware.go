package authentication

import (
	"fmt"
	"time"
	"net/http"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(publicKey *rsa.PublicKey) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		
	}
}
