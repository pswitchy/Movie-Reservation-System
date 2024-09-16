package handlers

import (
    "encoding/json"
    "net/http"
    "movie-reservation-system/models"
    "movie-reservation-system/db"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
    var movie models.Movie
    json.NewDecoder(r.Body).Decode(&movie)
    db.DB.Create(&movie)
    json.NewEncoder(w).Encode(movie)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
    var movies []models.Movie
    db.DB.Preload("Showtimes").Find(&movies)
    json.NewEncoder(w).Encode(movies)
}
