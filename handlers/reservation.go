package handlers

import (
    "encoding/json"
    "net/http"
    "movie-reservation-system/models"
    "movie-reservation-system/db"
)

func ReserveSeat(w http.ResponseWriter, r *http.Request) {
    var reservation models.Reservation
    json.NewDecoder(r.Body).Decode(&reservation)

    var count int64
    db.DB.Model(&models.Reservation{}).Where("showtime_id = ? AND seat_number = ?", reservation.ShowtimeID, reservation.SeatNumber).Count(&count)
    if count > 0 {
        http.Error(w, "Seat already reserved", http.StatusConflict)
        return
    }

    db.DB.Create(&reservation)
    json.NewEncoder(w).Encode(reservation)
}

func GetReservations(w http.ResponseWriter, r *http.Request) {
    var reservations []models.Reservation
    db.DB.Find(&reservations)
    json.NewEncoder(w).Encode(reservations)
}

func GetRevenue(w http.ResponseWriter, r *http.Request) {
    var revenue float64
    db.DB.Model(&models.Reservation{}).Select("SUM(price)").Row().Scan(&revenue)
    json.NewEncoder(w).Encode(map[string]float64{"revenue": revenue})
}
