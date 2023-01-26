package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
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


	fmt.Print("Starting")
	log.fatal(http.ListenAndServe(":8000",r))

	movies=append(movies, Movie{ID : "1", isbn: "Hi", Title: "Anmol the hero", Director:&Director{firstname: "Anmol",lastname: "monga"}})
}

func getMovies(w http.ResponseWriter, r* http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r* http.Request){

	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)

	for index,item := range movies{
		if item.ID== params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r* http.Request){

	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)

	for _ ,item : = range movies{

		if item.ID == params["id"]{

			json.NewEncoder(w).Encode(item)
			return
		}

	}

}

func createMovies(w http.ResponseWriter, r* http.Request){

	w.Header().Set("content-type","application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.(rand.Intn(1000000))

	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovies(w http.ResponseWriter, r* http.Request){

	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)

	for index,item :=range movies{
		
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}