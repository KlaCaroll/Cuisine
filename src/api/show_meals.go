package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func (s Service) ShowMeals(w http.ResponseWriter, r *http.Request) {
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

	var mealIDs = make([]int64, 0, len(meals))
	var mealsByID = make(map[int64]Meal, len(meals))
	for _, meal := range meals {
		mealIDs = append(mealIDs, meal.ID)
	}

	query, args, _ := sqlx.In(`
		select mr.meal_id, id, name
		from recipe as r
		join meal_recipe as mr on r.id = mr.recipe_id
		where mr.meal_id in (?)
	`, mealIDs)

	var res []struct{
		MealID int64 `db:"meal_id"`
		Recipe
	}
	err = s.DB.Select(&res, query, args...)
	_ = err

	for _, r := range res {
		meal := mealsByID[r.MealID]
		meal.Recipes = append(meal.Recipes, r.Recipe)
		mealsByID[r.MealID] = meal
	}

	meals = make([]Meal, 0, len(meals))
	for _, m := range mealsByID {
		meals = append(meals, m)
	}

	write(w, meals)
}