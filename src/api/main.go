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

type any = interface{}

func main() {
	var (
		addr string
		dsn  string // Data Source Name
	)
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "addr to listen on")
	flag.StringVar(&dsn, "dsn", "database.sqlite", "path to the database to use")
	flag.Parse()

	// INITIALIZE THE DATABASE CONNEXION
	log.Println("opening connection to", dsn)
	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		log.Println("opening connection", err)
		return
	}
	log.Println("opened connection")
	defer db.Close()

	var mux = http.NewServeMux()
	
	// CREATE MEAL
	mux.HandleFunc("/createMeal", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			PlannedAt time.Time `json:"planned_at"`
			Guests int64 `json:"guests"`
		}

		err := read(r, &input)
		if err!= nil {
			log.Println("parsing input", err)
			writeError(w,"input_error", err)
			return
		}

		res, err := db.Exec(`
			INSERT INTO meal (planned_at, guests)
			VALUES (?, ?)
		`, input.PlannedAt, input.Guests)
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "database_error", err)
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

	// SHOW MEAL
	mux.HandleFunc("/showMeals", func(w http.ResponseWriter, r *http.Request) {
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

	// DELETE MEAL
	mux.HandleFunc("/deleteMeal", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID int64 `json:="id"`
		}

		err := read(r, &input)
		if err!= nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		res, err := db.Exec(`
			DELETE FROM meal
			WHERE id=?
		`, input.ID)
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "database_error", err)
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
	// UPDATE MEAL refaire les champs
	mux.HandleFunc("/updateMeal", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID int64 `json:"id"`
			PlannedAt time.Time `json:"planned_at"`
			Guests int64 `json:"guests"`
		}

		err := read(r, &input)
		if err!= nil {
			log.Println("parsing input", err)
			writeError(w, "input_error", err)
			return
		}

		res, err := db.Exec(`
			UPDATE meal
			SET planned_at = ?, guests = ?
			WHERE id = ?
		`,input.PlannedAt, input.Guests, input.ID)
		if err != nil {
			log.Println("querying database", err)
			writeError(w, "database_error", err)
			return
		}

		log.Println("Meal updated", res)
	})
	
	// Add recipe (avantage / inconvenients / possible ou pas ...)

	// Create and populate the router.
	mux.HandleFunc("/computeShoppingList", func(w http.ResponseWriter, r *http.Request) {
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

func read(r *http.Request, payload any) (err error) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, payload)
	if err != nil {
		return err
	}
	return nil		
}