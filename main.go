package main

import (
	"fmt"
	"log" //logs out error
	"encoding/json" // I want to encode my data into json for sending it to postman
	"math/rand" // To create new movie we have to create new ID so we will generate that randomly 
	"net/http" //This allows us to create a server in the go
	"strconv" // The ID that we will create using math.random will be an integer hence this will convert that to string 
	"github.com/gorilla/mux" //Library , provides advanced routing features, middleware support, and optimized performance
)

//so basically I have created a struct of movie in which I am initializing these things and all of them have string as a data type.


type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"` // I am using pointer here because, I will create another struct of director and I want to associate Director with movie.
}

//Now Director is a pointer in the struct movie here I have defined the director as struct. 
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie // I have created a slice, which is a dynamic array, also it's type is Movie which is the struct we have defined above

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_= json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
		
	}
}



func main(){ //every go code must have func which handles everything. 
	r := mux.NewRouter() // this helps creating new routers we are using gorilla mux for it 

	// Below two lines append two movie types in the slice movies we have created 

	movies = append(movies, Movie{ID: "1", Isbn: "538227", Title: "Movie One", Director: &Director{Firstname: "Jhon", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "345789", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	//below I have created different functions to handle different routes and also wrote there Methods.

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // this line starts the server at the given port using net/http package.
}
