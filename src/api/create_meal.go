package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PlannedAt time.Time `json:"planned_at"`
		Guests int64 `json:"guests"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w,"input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO meal (planned_at, guests)
		VALUES (?, ?)
	`, input.PlannedAt, input.Guests)
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