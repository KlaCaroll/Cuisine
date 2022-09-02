package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) ShowMeal(w http.ResponseWriter, r *http.Request) {
	var input struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	var meal []struct{
		ID int64 `db:"id" json:"id"`
		PlannedAt time.Time `db:"planned_at" json:"planned_at"`
		Guests uint `db:"guests" json:"guests"`
	}

	err = s.DB.Select(&meal, `
		SELECT * FROM meal
		WHERE planned_at BETWEEN ? AND ?
	`, input.From, input.To)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, meal)
}