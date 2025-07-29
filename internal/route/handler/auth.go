package handler

import (
	"net/http"
)


func AuthenticationHandler(privateKeyPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
