package main

import (
//	"fmt"
	"log"
	"net/http"
)

func fetch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("connexion ok"));
}

func main() {
	http.HandleFunc("/connect", fetch)


	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}