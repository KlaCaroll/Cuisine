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
	quantity float64
)

func createMeals(w http.ResponseWriter, r *http.Request) {	
	//fmt.Println("what eat ? :")
	//fmt.Scanf("%v", &name)
	//fmt.Println("and when ? :")
	//fmt.Scanf("%v", &date)
	//fmt.Println("so, we will eat", name, "at" , date)


	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("REPLACE INTO meal(id, name, planned_at) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert datas")
	res, err := stmt.Exec(1306221, "patate", "2022-06-13")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	
	rows, err := db.Query("SELECT name, planned_at FROM meal")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		rows.Scan(&name, &planned_at)
		fmt.Println(name, planned_at)
	}


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
	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	rows, err := db.Query("select name from meal where id = '2205171'; ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		rows.Scan(&name)
		fmt.Println(name)
	}



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