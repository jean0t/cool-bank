package route

import (
	"net/http"

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
	var (
		function http.HandlerFunc
		ok bool
	)

	function, ok = h.handlers[path]
	if !ok {
		var not_found = func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not Found", http.StatusNotFound)	
		}

		return not_found
	}

	return function
}



//=======================================================================| Router

func registerHandlers(publicKeyPath, privateKeyPath string) *Handler {
	var handlers *Handler = &Handler{handlers: map[string]http.HandlerFunc{}}

	handlers.AddHandler("/account/", authentication.AuthMiddleware(publicKeyPath)(handler.AccountHandler))
	handlers.AddHandler("/admin/", authentication.AuthMiddleware(publicKeyPath)(authentication.ManagerOnly(handler.AdminHandler)))
	handlers.AddHandler("/auth/", handler.AuthenticationHandler(privateKeyPath))

	return handlers
}


func CreateRouter(publicKeyPath, privateKeyPath string) *http.ServeMux {
	var (
		handlers *Handler = registerHandlers(publicKeyPath, privateKeyPath)
		router *http.ServeMux = http.NewServeMux() 
	)

	// account routes
	router.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/account/", http.StatusMovedPermanently)
	})
	router.Handle("/account/", handlers.GetHandler("/account/"))

	// manager routes
	router.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/", http.StatusMovedPermanently)
	})
	router.Handle("/admin/", handlers.GetHandler("/admin/"))


	// authentication routes
	router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/auth/", http.StatusMovedPermanently)
	})
	router.Handle("/auth/", handlers.GetHandler("/auth/"))


	return router
}
