package handlers

import (
    "encoding/json"
    "net/http"
    "movie-reservation-system/models"
    "movie-reservation-system/db"
)

func CreateShowtime(w http.ResponseWriter, r *http.Request) {
    var showtime models.Showtime
    json.NewDecoder(r.Body).Decode(&showtime)
    db.DB.Create(&showtime)
    json.NewEncoder(w).Encode(showtime)
}
