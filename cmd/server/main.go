package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Diego0Leme/theater/pkg/http/rest"
	"github.com/Diego0Leme/theater/pkg/management"
	managementSql "github.com/Diego0Leme/theater/pkg/storage/sql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "testdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repos, _ := managementSql.NewRepository(db)
	service, _ := management.NewService(repos)

	router := rest.CreateRouter(service)
	http.ListenAndServe(":8080", router)
}
