package mem

import "github.com/go-hello/pkg/management"

type movieMemoryRepository struct{}

var movies []*management.Movie = make([]*management.Movie, 0)

func NewRepository() (*movieMemoryRepository, error) {
	return &movieMemoryRepository{}, nil
}

func (r *movieMemoryRepository) AddMovie(m *management.Movie) error{ 
	movies = append(movies, m)
	return nil
}

func (r *movieMemoryRepository) GetMovies() ([]*management.Movie, error){
	return movies, nil
}
