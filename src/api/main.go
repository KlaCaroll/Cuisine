package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	_ "github.com/mattn/go-sqlite3"
)

type any = interface{}

type Service struct {
	DB *sqlx.DB
}	

// Add recipe (avantage / inconvenients / possible ou pas ...)

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

	s := Service {
		DB: db,
	}

	// add recipe to meal + delete + 
	// dans create et update ajouter les tables pivo

	var mux = http.NewServeMux()

	mux.HandleFunc("/listRecipes", s.ListRecipes)
	mux.HandleFunc("/showRecipe", s.ShowRecipe)
	mux.HandleFunc("/createRecipe", s.CreateRecipe)
	mux.HandleFunc("/deleteRecipe", s.DeleteRecipe)

	mux.HandleFunc("/createMeal", s.CreateMeal)
	mux.HandleFunc("/showMeals", s.ShowMeals)
	mux.HandleFunc("/deleteMeal", s.DeleteMeal) 
	mux.HandleFunc("/updateMeal", s.UpdateMeal)
	mux.HandleFunc("/computeShoppingList", s.ComputeShoppingList)

	//mux.HandleFunc("/addRecipeToMeal", s.AddRecipeToMeal)
	//mux.HandleFunc("/deleteRecipeFromMeal", s.deleteRecipeFromMeal)

	// Start the HTTP server.
	handler := cors.Default().Handler(mux)
	// CORS => handler
	var srv = &http.Server{
		Addr:    addr,
		Handler: handler,
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