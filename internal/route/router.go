package route

import (
	"net/http"
	"fmt"
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


func CreateRouter() *http.ServeMux {
	var (
		handler *Handler = registerHandlers()
		router *http.ServeMux = http.NewServeMux() 
	)

	router.HandleFunc("/", handler.GetHandler("/"))

	return router
}
