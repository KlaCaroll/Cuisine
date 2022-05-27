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
	w.Write([]byte("connexion database"));
}

func main() {
	os.Remove("data/database.db")

	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//rows, err := db.Query("SELECT name from meal")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()

	http.HandleFunc("/connect", fetch)


	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}