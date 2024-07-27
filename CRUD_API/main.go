package main

import(
	"fmt"
	"log"
	"encoding/json"   // to encode data when sending it to pasotman
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type movie struct{
	ID string `json:"id` // json has to be in smallcap so we can encode and decode it
	Isbn string `json:"isbn`
	Title string `json:title`
	Diector *Director `json:director`
}

type Director struct{
	Firstname string `json:"firstname`
	Lastname string `json:"lastname`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Tuype","application/json")
	json.NewEncoder(w).Encoder(movies)
}

func deleteMovies(w http.ResponceWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		
		if item.ID == params["id"]{
			movie = append(movie[:index],movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponceWriter, r *http.request){
	w.Header().Set("Content-Type", "application/json")
	params :=  ,ux.Varas(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func craeteMovie(w http.ResponceWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).decode(&movie)
	movie.ID = strconv(random.Intn(10000000))
	movies = append(movies, movie) 
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponceWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range moview{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

 func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{Id; "1", Isbn:"438227", Title: "Movie One", Director: &Director{Fisrstname:"John", Lastname:"Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn:"45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname:"Smith"}})
	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovies).Method("GET")
	r.HandleFunc("/movies", craeteMovie).Method("POST")
	r.Handlefunc("/movies/{id}", updateMovie).Method("PUT")
	r.handleFunc("/movies{id}",deleteMovie).Method("DELETE")

	fmt.Printf("Starting Server at 8000\n")
	log.Fatal(htp.ListenAndServer(":8000", r))
 }