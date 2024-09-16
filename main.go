package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "movie-reservation-system/db"
    "movie-reservation-system/handlers"
    "movie-reservation-system/middleware"
)

func main() {
    db.InitDB()
    r := mux.NewRouter()

    r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
    r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

    r.Handle("/movies", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateMovie))).Methods("POST")
    r.Handle("/movies", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetMovies))).Methods("GET")
    r.Handle("/showtimes", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateShowtime))).Methods("POST")
    r.Handle("/reservations", middleware.AuthMiddleware(http.HandlerFunc(handlers.ReserveSeat))).Methods("POST")
    r.Handle("/admin/reservations", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetReservations))).Methods("GET")
    r.Handle("/admin/revenue", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetRevenue))).Methods("GET")

    http.ListenAndServe(":8080", r)
}
