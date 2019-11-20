package main

import (
	"net/http"
	"github.com/Diego0Leme/theater/pkg/http/rest"
	"github.com/Diego0Leme/theater/pkg/storage/mem"
)

func main() {
	memRepos, _ := mem.NewRepository()
	
	router := rest.CreateRouter(memRepos)
	http.ListenAndServe(":8080", router)
}
