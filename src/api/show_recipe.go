package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) ShowRecipe(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RecipeID int64 `json:recipe_ID`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	var recipeFood []struct{
		Recipe_name string `db:"recipe_name" json:"recipe_name"`
		Food_name string `db:"food_name" json:"food_name"`
		Quantity float64 `db:"quantity" json:"quantity"`

	}

	err = s.DB.Select(&recipeFood, `
		SELECT quantity, food.name AS food_name, recipe.name AS recipe_name
		FROM recipe_food
		JOIN food ON food.id = recipe_food.food_id
		JOIN recipe ON recipe.id = recipe_food.recipe_id
		WHERE recipe_ID=?
	`, input.RecipeID)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, recipeFood)
}