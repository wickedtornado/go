package main

import (
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *director `json:"director"`
}

type Director struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updatetMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
}
