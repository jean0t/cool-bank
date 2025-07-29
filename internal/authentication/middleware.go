package authentication

import (
	"strings"
	"context"
	"net/http"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(publicKey *rsa.PublicKey) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var (
				authHeader, tokenString string
				claims *Claims
				ok bool
				ctx context.Context
				token *jwt.Token
				err error

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
