package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *director `json: "director"`
}

type director struct {
	FirstName string `json: "firstname`
	LastName  string `json: "lastname"`
}

var movies []movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &director{FirstName: "Guillermo", LastName: "del Torro"}})
	movies = append(movies, movie{ID: "2", Isbn: "89348", Title: "Another Movie", Director: &director{FirstName: "Jack", LastName: "Hollande"}})
	movies = append(movies, movie{ID: "3", Isbn: "981753", Title: "Jack Spanker", Director: &director{FirstName: "John", LastName: "Smith"}})
	movies = append(movies, movie{ID: "4", Isbn: "785614", Title: "Title", Director: &director{FirstName: "Ivan", LastName: "the Horrible"}})

	// Reads all the movies
	router.HandleFunc("/movies", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(movies)
	}).Methods("GET")

	// Reads one specific movie
	router.HandleFunc("/movies/{id}", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		parameters := mux.Vars(request)

		for _, movie := range movies {
			if movie.ID == parameters["id"] {
				json.NewEncoder(response).Encode(movie)
				return
			}
		}
	}).Methods("GET")

	// Creates a new movie
	router.HandleFunc("/movies", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		var new_movie movie
		_ = json.NewDecoder(request.Body).Decode(&new_movie) // Why are we doing this?
		new_movie.ID = strconv.Itoa(rand.Intn(100_000_000))
		movies = append(movies, new_movie)
		json.NewEncoder(response).Encode(new_movie)
	}).Methods("POST")

	// Modifies an existing movie
	router.HandleFunc("/movies/{id}", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		parameters := mux.Vars(request)

		for _, movie := range movies {
			if movie.ID == parameters["id"] {
				_ = json.NewDecoder(request.Body).Decode(&movie)

				movie.Isbn = parameters["isbn"]
				movie.Title = parameters["title"]
				movie.Director.FirstName = parameters["firstname"]
				movie.Director.LastName = parameters["lastname"]

				json.NewEncoder(response).Encode(movie)

				return
			}
		}
	}).Methods("PUT")

	// Deletes an existing movie
	router.HandleFunc("/movies/{id}", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		parameters := mux.Vars(request)

		for index, movie := range movies {
			if movie.ID == parameters["id"] {
				// Delete the movie with the particular index
				movies = append(movies[:index], movies[index+1:]...)
				break
			}
		}

		json.NewEncoder(response).Encode(movies)
	}).Methods("DELETE")

	fmt.Printf("Starting server on port 80\n")
	log.Fatal(http.ListenAndServe(":80", router))
}
