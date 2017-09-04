package handler

import (
	"net/http"
	"io"
	"strings"
	"database/sql"
	"log"
)

type Handler struct {
	Db *sql.DB
}

func (a Handler) List(w http.ResponseWriter, req *http.Request) {
	results, err := a.Db.Query("Select title from todos")
	if err != nil {
		log.Fatal(err)
	}
	defer results.Close()
	for results.Next() {
		var title string
		results.Scan(&title)
		io.WriteString(w, title+"\n")
	}
}

func (a Handler) Update(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.RawQuery, "id=")
	io.WriteString(w, id)

	//create new element
	if id == "" {

	} else {

	}
}
