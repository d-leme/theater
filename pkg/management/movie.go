package management

var Categories = map[string]string{
	"Action":"action"}

type Movie struct {
	Id       int32
	Name     string
	Category string
}

type MovieRepository interface {
	AddMovie(movie *Movie) error
	GetMovies() ([]*Movie, error)
}
