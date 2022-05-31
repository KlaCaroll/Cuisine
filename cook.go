package main

import (
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3" 
	"log"
	"net/http"
	//"os"
)

func fetch(w http.ResponseWriter, r *http.Request) {
	//os.Remove("data/database.db")
	var ingredients string
	var name string
	var quantity float64

	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
		fmt.Println("connexion failed")
	}
	w.Write([]byte("connexion database"));
	defer db.Close()

	rows, err := db.Query("select food_id as ingredients, food.name as name, quantity from recipe_food, food where food_id = food.id and recipe_id = '121191714519'; ")
	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ingredients, &name, &quantity)
		fmt.Println(ingredients, name, quantity)
	}
}

func main() {


	http.HandleFunc("/connect", fetch)


	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}