package handler

import (
	"net/http"
)


func AuthenticationHandler(privateKeyPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
			case "/auth/login", "/auth/login/":
				http.Error(w, "Not implemented", http.StatusNoContent)

			case "/auth/register", "/auth/register/":
				http.Error(w, "Not implemented", http.StatusNoContent)

			case "/auth/":
				http.Error(w, "Bad Request, Try /auth/login or /auth/register", http.StatusBadRequest)

			default:
				http.Error(w, "Bad Request", http.StatusBadRequest)
		}

	}
}
