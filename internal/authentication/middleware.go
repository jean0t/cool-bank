package authentication

import (
	"fmt"
	"strings"
	"context"
	"net/http"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(publicKeyPath string) func(http.HandlerFunc) http.HandlerFunc {
	var (
		publicKey *rsa.PublicKey
		err error
	)

	publicKey, err = loadPublicKey(publicKeyPath)
	if err != nil {
		fmt.Println("[*] error loading public key for rsa")
	}
	
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var (
				authHeader, tokenString string
				claims *Claims
				ok bool
				ctx context.Context
				token *jwt.Token
			)

			authHeader = r.Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}

			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			token, err = jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error){
				return publicKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			claims, ok = token.Claims.(*Claims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(r.Context(), "userClaims", claims)
			next(w, r.WithContext(ctx))
		}
	}
}


func ManagerOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ok bool
			claims *Claims
		)

		claims, ok = r.Context().Value("userClaims").(*Claims)
		if !ok || claims.Role != "manager" {
			http.Error(w, "Forbidden: Only Managers", http.StatusForbidden)
			return
		}

		next(w, r)
	}	
}
