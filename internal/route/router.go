package route

import (
	"net/http"
	"crypto/rsa"
	
	"fmt"

	"github.com/jean0t/cool-bank/internal/authentication"
	"github.com/jean0t/cool-bank/internal/route/handler"
)



//=======================================================================| Handlers

type Handler struct {
	handlers map[string]http.HandlerFunc
}

func (h *Handler) AddHandler(path string, handlerFunc http.HandlerFunc) {
	h.handlers[path] = handlerFunc
}

func (h *Handler) GetHandler(path string) http.HandlerFunc {
	return h.handlers[path]
}



//=======================================================================| Router

func registerHandlers() *Handler {
	var handler *Handler = &Handler{handlers: map[string]http.HandlerFunc{}}

	handler.AddHandler("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Apenas um teste</h1>")
	})
	return handler
}


func CreateRouter(publicKeyPath, privateKeyPath string) *http.ServeMux {
	var (
		handler *Handler = registerHandlers()
		router *http.ServeMux = http.NewServeMux() 
	)

	// account routes
	router.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/account/", http.StatusMovedPermanently)
	})
	router.Handle("/account/", authentication.AuthMiddleware(publicKey)(handler.AccountHandler))

	// manager routes
	router.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/", http.StatusMovedPermanently)
	})
	router.Handle("/admin/", authentication.AuthMiddleware(publicKey)(authentication.ManagerOnly(handler.AdminHandler)))


	// authentication routes
	router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/auth/", http.StatusMovedPermanently)
	})
	router.Handle("/auth/", handler.AuthenticationHandler(privateKeyPath))


	return router
}
