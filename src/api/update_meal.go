package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) UpdateMeal(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int64 `json:"id"`
		PlannedAt time.Time `json:"planned_at"`
		Guests int64 `json:"guests"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		UPDATE meal
		SET planned_at = ?, guests = ?
		WHERE id = ?
	`,input.PlannedAt, input.Guests, input.ID)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	log.Println("Meal updated", res)
}