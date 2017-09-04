package main

import (
	"net/http"
	"todoapp/handler"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://todoadmin:superadmin@go-meetup-postgres.cgh71ig5ptkc.eu-central-1.rds.amazonaws.com:5432/todo")

	if err != nil {
		log.Fatal(err)
	}

	h := handler.Handler{Db:db}

	http.HandleFunc("/list", h.List)
	http.HandleFunc("/update", h.Update)
	http.ListenAndServe(":8080", nil)
}
