package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) listRecipes(w http.ResponseWriter, r *http.Request) {

	var recipe []struct{
		ID int64 `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}

	err := s.DB.Select(&recipe, `
		SELECT * FROM recipe
	`)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, recipe)
}