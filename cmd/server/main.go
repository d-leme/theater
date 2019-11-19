package main

import (
	"net/http"
	"github.com/go-hello/pkg/http/rest"
	"github.com/go-hello/pkg/storage/mem"
)

func main() {
	memRepos, _ := mem.NewRepository()
	
	router := rest.CreateRouter(memRepos)
	http.ListenAndServe(":8080", router)
}
