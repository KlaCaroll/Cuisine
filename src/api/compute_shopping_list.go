package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) ComputeShoppingList(w http.ResponseWriter, r *http.Request) {
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

	var items []struct {
		Name     string  `db:"name" json:"name"`
		Quantity float64 `db:"quantity" json:"quantity"`
	}
	
	err = s.DB.Select(&items, `
		SELECT f.name AS name, rf.quantity * m.guests AS quantity
		FROM meal m
		JOIN meal_recipe mr ON m.id = mr.meal_id
		JOIN recipe_food rf ON mr.recipe_id = rf.recipe_id
		JOIN food f ON rf.food_id = f.id
		WHERE m.planned_at BETWEEN ? AND ?
	`, input.From, input.To)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, items)
}