package main

import (
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3" 
	"log"
	"net/http"
	//"os"
)

var servive struct {
	db *sql.DB
}

var (
	id int
	name string
	date string
	sub_type string
	pers int
	quantity float64
)

func createMeals(w http.ResponseWriter, r *http.Request) {
	//db, err := sql.Open("sqlite3", "data/database.db")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	fmt.Println("what eat ?")
	fmt.Scanf("%v", &name)
	fmt.Println("we will eat : ", name)
	fmt.Println("when ?")
	fmt.Scanf("%v", &date)
	fmt.Println("at", date)
	fmt.Println("so, we will eat", name, "at" , date)
}
func fetchMeals(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetchMeals")
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
	fmt.Println("list")
}

func main() {

	//rows, err := db.Query("select food.name as name, quantity from recipe_food, food where food_id = food.id and recipe_id = '318151721513151419952118'; ")
	//if err != nil {
	//	log.Fatal(err)
	//	fmt.Println("error")
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	rows.Scan(&name, &quantity)
	//	fmt.Println(name, quantity)
	//}

	http.HandleFunc("/createMeals", createMeals)
	http.HandleFunc("/fetchMeals", fetchMeals)
	http.HandleFunc("/createRecipe", createRecipe)
	http.HandleFunc("/fetchRecipe", fetchRecipe)
	http.HandleFunc("/createFood", createFood)
	http.HandleFunc("/fetchFood", fetchFood)
	http.HandleFunc("/list", list)

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}