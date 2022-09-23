package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) ListRecipes(w http.ResponseWriter, r *http.Request) {
	var recipes []struct{
		ID int64 `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}

	err := s.DB.Select(&recipes, `
		SELECT id, name FROM recipe
	`)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, recipes)
}