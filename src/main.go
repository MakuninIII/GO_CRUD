package main

import (
	"MusicBands/src/handlers"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS bands (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        country TEXT NOT NULL,
        debut_year INTEGER
        );
    `)

	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	handlers.Initialize(db)

	r := mux.NewRouter()
    r.HandleFunc("/token", handlers.GetToken).Methods("GET")
	r.Handle("/bands", handlers.JWTMiddleware(http.HandlerFunc(handlers.CreateBand))).Methods("POST")
	r.Handle("/bands", handlers.JWTMiddleware(http.HandlerFunc(handlers.GetBands))).Methods("GET")
	r.Handle("/bands/{id}", handlers.JWTMiddleware(http.HandlerFunc(handlers.UpdateBand))).Methods("PUT")
	r.Handle("/bands/{id}", handlers.JWTMiddleware(http.HandlerFunc(handlers.DeleteBand))).Methods("DELETE")

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
