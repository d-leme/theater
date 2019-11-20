package management

import (
	"errors"
	"github.com/satori/go.uuid"
)

type ( 
	Category string

	Movie struct {
		Id       uuid.UUID
		Name     string
		Category Category
	}
)

var UnkownCategoryError = errors.New("Unkown category")

const (
	Action Category = "action"
)

var Categories = map[string]Category{
	"Action": Action,
}

func NewCategory(c string) (Category, error){
	category := Categories[c]
	if category == ""{
		return "", UnkownCategoryError
	}

	return category, nil
}


type MovieRepository interface {
	AddMovie(movie *Movie) error
	GetMovies() ([]*Movie, error)
}
