package handler

import (
	"net/http"

	"github.com/jean0t/cool-bank/internal/controller"
)


func AuthenticationHandler(privateKeyPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
			case "/auth/login", "/auth/login/":
				return controller.Login(w, r)

			case "/auth/register", "/auth/register/":
				return controller.Register(w, r)

			case "/auth/":
				http.Error(w, "Bad Request, Try /auth/login or /auth/register", http.StatusBadRequest)

			default:
				http.Error(w, "Bad Request", http.StatusBadRequest)
		}

	}
}
