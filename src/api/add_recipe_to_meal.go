package main

import (
	"log"
	"net/http"
//	"time"
//	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) AddRecipeToMeal(w http.ResponseWriter, r *http.Request) {
	var input struct{
		Meal_id int64 `json:"meal_id"`
		Recipe_id int64 `json:"recipe_id"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	_ , err = s.DB.Exec(`
		INSERT INTO meal_recipe (meal_id, recipe_id)
		VALUES (?,?)
	`, input.Meal_id, input.Recipe_id)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, "ok")
}