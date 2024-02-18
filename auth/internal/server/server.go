package server

import (
	"cards-service/auth/internal/service/auth"
	jwtservice "cards-service/auth/internal/service/jwt_service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router       *chi.Mux
	authServices map[string]auth.Auth
	jwtService   *jwtservice.JwtService
}

func NewServer(jwtService *jwtservice.JwtService, services ...auth.Auth) *Server {
	r := chi.NewRouter()
	serviceMap := make(map[string]auth.Auth)
	for _, s := range services {
		serviceMap[s.Key()] = s
	}
	server := Server{router: r, authServices: serviceMap, jwtService: jwtService}
	server.initRouter()
	return &server
}

func (s *Server) GetRouter() *chi.Mux {
	return s.router
}
