package rest

import (
	"github.com/go-hello/pkg/management"
	"github.com/go-chi/chi"
)

func CreateRouter(mRepos management.MovieRepository) chi.Router {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router){
		h := managementHandler{repo:mRepos}
		r.Mount("/v1", h.router())
	})
	
	return r
}


