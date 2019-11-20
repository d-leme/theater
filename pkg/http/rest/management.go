package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Diego0Leme/theater/pkg/management"
)

type managementHandler struct{
	repo management.MovieRepository
}

func (m *managementHandler) router() chi.Router{
	r := chi.NewRouter()

	r.Route("/management", func(r chi.Router){
		r.Get("/movies", m.getMovies)
		r.Post("/movies", m.createMovie)
	})

	return r
}

func (m *managementHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, _ := m.repo.GetMovies()

	json.NewEncoder(w).Encode(&movies)
	w.WriteHeader(http.StatusOK)
}

func (m *managementHandler) createMovie(w http.ResponseWriter, r *http.Request) {

	var request struct{
		Name string `json:name`
		Category string `json:category`
	}	

	movies, _ := m.repo.GetMovies()

	id := int32(len(movies) + 1)
	category := management.Categories[request.Category]

	json.NewDecoder(r.Body).Decode(&request)
	movie := management.Movie{
		Id: id,
		Name:request.Name,
		Category: category,
	}

	err := m.repo.AddMovie(&movie)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}