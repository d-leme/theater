package management

import "github.com/satori/go.uuid"

type Service interface{
	CreateMovie(n string, c string) error
	GetMovies() ([]*Movie, error)
}

type service struct{
	m MovieRepository
}

func NewService(movieRepo MovieRepository) (*service, error){
	return &service{m:movieRepo}, nil
}

func (s *service) CreateMovie(n string, c string) error{

	category, err := NewCategory(c)
	if err != nil{
		return err
	}

	id, _ := uuid.NewV4()
	movie := Movie{
		Id:id,
		Name:n,
		Category: category,
	}

	if err := s.m.AddMovie(&movie); err != nil{
		return err
	}

	return nil
}

func (s *service) GetMovies() ([]*Movie, error){

	movies, err := s.m.GetMovies()
	if err != nil{
		return nil, err
	}

	return movies, nil
}