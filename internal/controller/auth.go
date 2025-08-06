package controller

import (
	"net/http"
	"encoding/json"

	"github.com/jean0t/cool-bank/internal"
)


type Info struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}


func Login(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			jwt string
			err error
		)
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are accepted", http.StatusMethodNotAllowed)
			return
		}
		
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		
		var info Info
		err = json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			fmt.Println("[!] Error decoding JSON")
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		jwt, err = CreateJWT(r.Name, r.Role, 1 * time.Hour, cli.PrivateKeyPath)
		w.Header().Set("Authorization", "Bearer " + jwt)
		w.WriteHeader(http.StatusOK)
		var response = map[string]string{"message": "Token Generated"}
		json.NewEncoder(w).Encode(response)
	}
}


func Register(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
