package main

import (
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3" 
	"log"
	"net/http"
	//"os"
)

var (
	id int
	name string
	planned_at string
	sub_type string
	pers int
	quantity_r float64
	food string
	recipe_id int
	recipe_name string
)

func createMeal(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("REPLACE INTO meal(id, name, planned_at) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(1306221, "patate", "2022-06-13")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp, "meal created")
	//return os.WriteFile("data/" + seed + ".sql", 0600)
}

func deleteMeal(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Exec("DELETE FROM meal where name = 'patate';")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stmt,"meal deleted")
	http.Redirect(w, r, "/fetchMeal", http.StatusFound)
}

func fetchMeal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MEALS :")

	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	rows, err := db.Query("SELECT name, planned_at FROM meal; ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&name, &planned_at)
		fmt.Println(name, planned_at)
	}
}
func createRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createRecipe")
}

func fetchRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetchRecipe")
}
func createFood(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createFood")
}
func fetchFood(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetchFood")
}
func list(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT recipe_name, food_name, quantity_r FROM recipe, recipe_food, food WHERE recipe.id = recipe_food.recipe_id AND recipe_food.food_id = food.id AND recipe_name = 'quiche_lorraine';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&recipe_name, &food, &quantity_r,)
		fmt.Println(recipe_name, food, quantity_r)
	}
}

func main() {
	http.HandleFunc("/createMeal", createMeal)
	http.HandleFunc("/fetchMeal", fetchMeal)
	http.HandleFunc("/deleteMeal", deleteMeal)
	http.HandleFunc("/createRecipe", createRecipe)
	http.HandleFunc("/fetchRecipe", fetchRecipe)
	http.HandleFunc("/createFood", createFood)
	http.HandleFunc("/fetchFood", fetchFood)
	http.HandleFunc("/list", list)

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}