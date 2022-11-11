package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"errors"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) ShowMeals(w http.ResponseWriter, r *http.Request) {
	var input struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	}

	err := read(r, &input)
	if err != nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	type Recipe struct{
		ID int64 `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}

	type Meal struct{
		ID int64 `db:"id" json:"id"`
		PlannedAt time.Time `db:"planned_at" json:"planned_at"`
		Guests uint `db:"guests" json:"guests"`
		Recipes []Recipe `json:"recipes"`
	}

	var meals []Meal

	err = s.DB.Select(&meals, `
		SELECT id, planned_at, guests
		FROM meal
		WHERE planned_at BETWEEN ? AND ?
	`, input.From, input.To)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	if len(meals) == 0 {
		write(w, meals)
		return
	}

	var mealIDs = make([]int64, 0, len(meals))
	var mealsByID = make(map[int64]Meal, len(meals))
	for _, meal := range meals {
		mealIDs = append(mealIDs, meal.ID)
		mealsByID[meal.ID] = meal
	}

	query, args, err := sqlx.In(`
		SELECT mr.meal_id, id, name
		FROM recipe as r
		JOIN meal_recipe as mr ON r.id = mr.recipe_id
		WHERE mr.meal_id in (?)
	`, mealIDs)
	if err !=nil {
		log.Println("building query", err)
		writeError(w,"query error", err)
	}

	var res []struct{
		MealID int64 `db:"meal_id"`
		Recipe
	}

	err = s.DB.Select(&res, query, args...)

	for _, r := range res {
		meal, ok := mealsByID[r.MealID]
		if !ok {
			log.Println("fail", r.MealID, mealsByID)
			writeError(w, "internal_error", errors.New("meal not found"))
			return
		}
		meal.Recipes = append(meal.Recipes, r.Recipe)
		mealsByID[r.MealID] = meal
	}


	fmt.Println(meals)
	meals = make([]Meal, 0, len(meals))
	for _, m := range mealsByID {
		meals = append(meals, m)
	}

	fmt.Println(meals)

	write(w, meals)
}