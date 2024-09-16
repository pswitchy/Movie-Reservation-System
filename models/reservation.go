package models

import "gorm.io/gorm"

type Reservation struct {
    gorm.Model
    UserID     uint
    ShowtimeID uint
    SeatNumber string
    Status     string
}

