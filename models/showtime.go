package models

import "gorm.io/gorm"

type Showtime struct {
    gorm.Model
    MovieID     uint
    StartTime   string
    EndTime     string
    Reservations []Reservation `gorm:"foreignKey:ShowtimeID"`
}



