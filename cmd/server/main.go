package main

import (
	"database/sql"
	"fmt"
	"os"
	"net/http"

	"github.com/Diego0Leme/theater/pkg/http/rest"
	"github.com/Diego0Leme/theater/pkg/management"
	managementSql "github.com/Diego0Leme/theater/pkg/storage/sql"
	_ "github.com/lib/pq"
)

func main() {
	var (
		host = os.Getenv("SQL_HOST")
		port = os.Getenv("SQL_PORT")
		user = os.Getenv("SQL_USER")
		password = os.Getenv("SQL_PASSWORD")
		dbname = os.Getenv("SQL_DBNAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
