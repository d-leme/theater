package main

import (
	"fmt"
	"github.com/Diego0Leme/theater/pkg/management"
	"github.com/Diego0Leme/theater/pkg/storage/mem"
)

func main() {
	memRepos, _ := mem.NewRepository()
	

	fstMovie := management.Movie{
		Id : 1,
		Name : "Joker",
		Category: "Action"}

	memRepos.AddMovie(&fstMovie)

	movies, _ := memRepos.GetMovies()

	for _, movie := range movies {
		fmt.Printf(movie.Name)
		fmt.Printf("%s",movie.Category)
	}
}