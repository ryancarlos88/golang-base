package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []WebHandler
	WebServerPort string
}

type WebHandler struct {
	method, path string
	handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]WebHandler, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path, method string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, WebHandler{handler: handler, method: method, path: path})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, x := range s.Handlers {
		s.Router.MethodFunc(x.method, x.path, x.handler)
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
