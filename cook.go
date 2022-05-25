package main

import (
//	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3" 
	"log"
	"net/http"
	"os"
)

func fetch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("connexion ok"));
}

func main() {
	os.Remove("data/database.db")

	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT food.name as ingredients, quantity from recipe_food, food where food_id = food.id and recipe_id = 121191714519;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	http.HandleFunc("/connect", fetch)


	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}