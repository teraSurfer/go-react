package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Movie struct {
	gorm.Model
	Name        string
	Director    string
	ReleaseDate time.Time
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&Movie{})
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin sslmode=disable")
	if err != nil {
		panic("Cound not connect to database")
	}
	defer db.Close()
	var movies []Movie
	db.Find(&movies)
	if len(movies) > 0 {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(movies)
	} else {
		http.Error(
			w,
			"No content found",
			http.StatusNoContent,
		)
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin sslmode=disable")
	if err != nil {
		panic("Cound not connect to database")
	}
	defer db.Close()
	id := mux.Vars(r)["id"]
	var movie Movie
	db.Where("id = ?", id).Find(&movie)
	if movie.ID != 0 {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(movie)
	} else {
		http.Error(
			w,
			"No content found",
			http.StatusNoContent,
		)
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin sslmode=disable")
	if err != nil {
		panic("Cound not connect to database")
	}
	defer db.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request",
			http.StatusBadRequest)
	}
	var movie Movie
	fmt.Println(r.Body)
	json.Unmarshal(body, &movie)
	// fmt.Println(movie)
	db.Create(&movie)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie created successfully")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin sslmode=disable")
	if err != nil {
		panic("Cound not connect to database")
	}
	defer db.Close()
	id := mux.Vars(r)["id"]
	var movie Movie
	db.Where("id = ?", id).Find(&movie)
	db.Delete(&movie)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Movie deleted successfully")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin sslmode=disable")
	if err != nil {
		panic("Cound not connect to database")
	}
	defer db.Close()
	id := mux.Vars(r)["id"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request",
			http.StatusBadRequest)
	}
	var movie Movie
	json.Unmarshal(body, &movie)
	db.Where("id = ?", id).Update(&movie)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Movie updated successfully")
}
