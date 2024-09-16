package models

import "gorm.io/gorm"

type Movie struct {
    gorm.Model
    Title       string
    Description string
    PosterImage string
    Genre       string
    Showtimes   []Showtime `gorm:"foreignKey:MovieID"`
}

