package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)
//ajouter create meal

type any = interface{}

func main() {
	var (
		addr string
		dsn  string // Data Source Name
	)
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "addr to listen on")
	flag.StringVar(&dsn, "dsn", "database.sqlite", "path to the database to use")
	flag.Parse()

	// Initialize the database connection.
	log.Println("opening connection to", dsn)
	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		log.Println("opening connection", err)
		return
	}
	log.Println("opened connection")
	defer db.Close()

	var mux = http.NewServeMux()
	
	// Create meal
	mux.HandleFunc("/createMeal", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			PlannedAt time.Time `json:"planned_at"`
			Guests int64 `json:"guests"`
		}

		raw, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		err = json.Unmarshal(raw, &input)
		if err != nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		res, err := db.Exec(`
			INSERT INTO meal (planned_at, guests)
			VALUES (?, ?)
		`, input.PlannedAt, input.Guests)
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "createMeal_error", err)
			return
		}

		var output struct{
			ID int64 `db:"id" json:"id"`
		}

		output.ID, err = res.LastInsertId()
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "database_error", err)
			return
		}

		write(w, output)
	})

	// Show Meal get ou list
	mux.HandleFunc("/showMeals", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			From time.Time `json:"from"`
			To   time.Time `json:"to"`
		}

		raw, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		err = json.Unmarshal(raw, &input)
		if err != nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		var meal []struct{
			ID int64 `db:"id" json:"id"`
			PlannedAt time.Time `db:"planned_at" json:"planned_at"`
			Guests uint `db:"guests" json:"guests"`
		}

		err = db.Select(&meal, `
			SELECT * FROM meal
			WHERE planned_at BETWEEN ? AND ?
		`, input.From, input.To)
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "database_error", err)
			return
		}

		write(w, meal)
	})

	// Delete Meal
	// Update Meal
	
	
	// Add recipe (avantage / inconvenients / possible ou pas ...)

	// Create and populate the router.
	mux.HandleFunc("/computeShoppingList", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			From time.Time `json:"from"`
			To   time.Time `json:"to"`
		}

		raw, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		err = json.Unmarshal(raw, &input)
		if err != nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		var items []struct {
			Name     string  `db:"name" json:"name"`
			Quantity float64 `db:"quantity" json:"quantity"`
		}
		err = db.Select(&items, `
			SELECT f.name AS name, rf.quantity * m.guests AS quantity
			FROM meal m
			JOIN meal_recipe mr ON m.id = mr.meal_id
			JOIN recipe_food rf ON mr.recipe_id = rf.recipe_id
			JOIN food f ON rf.food_id = f.id
			WHERE m.planned_at BETWEEN ? AND ?
		`, input.From, input.To)
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "database_error", err)
			return
		}

		write(w, items)
	})

	// Start the HTTP server.
	var srv = &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Println("listen on addr", addr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("listening", err)
		return
	}
}

func write(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	raw, _ := json.Marshal(payload)
	w.Write(raw)
}

type apiError struct{
	Code string `json:"code"`
	Err string `json:"err"`
}

func writeError(w http.ResponseWriter, code string, err error) {
	write(w, apiError{
		Code: code,
		Err: err.Error(),
	})
}