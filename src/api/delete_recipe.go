package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int64 `json:"id"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w,"input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		DELETE FROM recipe
		WHERE id=?
	`, input.ID)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}
 
	var output struct{
		ID int64 `db:"id" json:"id"`
	}

	output.ID, err = res.LastInsertId()
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, output)
}