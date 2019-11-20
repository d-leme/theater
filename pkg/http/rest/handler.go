package rest

import (
	"github.com/Diego0Leme/theater/pkg/management"
	"github.com/go-chi/chi"
)

func CreateRouter(s management.Service) chi.Router {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router){
		h := managementHandler{s: s}        
		r.Mount("/v1", h.router())
	})

	return r
}