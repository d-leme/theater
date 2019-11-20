package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Diego0Leme/theater/pkg/management"
	"github.com/go-chi/chi"
)

type managementHandler struct {
	s management.Service
}

func (m *managementHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/management", func(r chi.Router) {
		r.Get("/movies", m.getMovies)
		r.Post("/movies", m.createMovie)
	})

	return r
}

func (m *managementHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.s.GetMovies()

	if err != nil {
		encodeError(err, w)
		return
	}

	var response []Movie = make([]Movie, 0)

	for _, movie := range movies {
		response = append(response, Movie{
			Id:       movie.Id.String(),
			Name:     movie.Name,
			Category: string(movie.Category),
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&response)
}

func (m *managementHandler) createMovie(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}

	json.NewDecoder(r.Body).Decode(&request)

	if err := m.s.CreateMovie(request.Name, request.Category); err != nil {
		encodeError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type Movie struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func encodeError(err error, w http.ResponseWriter) {
	switch err {
	case management.UnkownCategoryError:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	response := map[string]interface{}{
		"message": err.Error(),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
