package sql

import (
	"database/sql"
	
	"github.com/satori/go.uuid"
	"github.com/Diego0Leme/theater/pkg/management"
)

type movieSQLRepository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) (*movieSQLRepository, error) {
	return &movieSQLRepository{db:db}, nil
}

func (r *movieSQLRepository) AddMovie(m *management.Movie) error{
	command := "INSERT INTO movies(\"id\", name, category) VALUES($1, $2, $3)"
	_, err := r.db.Exec(command, m.Id.String(), m.Name, string(m.Category))
	return err
}

func (r *movieSQLRepository) GetMovies() ([]*management.Movie, error){
	var movies []*management.Movie
	query := "SELECT * FROM movies"

	rows, err := r.db.Query(query)
	if err != nil{
		return nil, err
	}

	for rows.Next() {
		var (
			id string
			name string
			category string
		)

		rows.Scan(&id, &name, &category)
		
		movieUUID, err := uuid.FromString(id)
		if err != nil {
			return nil, err
		}
		
		movies = append(movies, &management.Movie{
			Id: movieUUID,
			Name: name,
			Category: management.Category(category),
		})
	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return movies, nil
}
