# Movie Reservation System

## Overview
The Movie Reservation System is a web application project that allows users to browse movies, view showtimes, and make reservations. The system is built using Go for the backend, PostgreSQL for the database.

## Features
- User Registration and Authentication
- Browse Movies and Showtimes
- Make and Manage Reservations
- Admin Panel for Managing Movies and Showtimes

## Technologies Used
- **Backend**: Go, GORM, JWT
- **Database**: PostgreSQL
- **Authentication**: JWT, bcrypt

## Installation

### Prerequisites
- Go 1.16+
- PostgreSQL 13+
- Node.js 14+
- npm 6+

### Backend Setup
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/movie-reservation-system.git
    cd movie-reservation-system
    ```
2. Set up PostgreSQL:
    - Create a new PostgreSQL database.
    - Update the connection string in `db/init.go` with your database credentials.
3. Install Go dependencies:
    ```sh
    go mod tidy
    ```
4. Run the backend server:
    ```sh
    go run main.go
    ```


## Usage

### Register and Login
- Register a new user by providing a username and password.
- Login with the registered credentials to access the system.

### Browse Movies
- View the list of available movies and their showtimes.
- Search for movies using the search functionality.

### Make Reservations
- Select a movie and showtime to make a reservation.
- View and manage your reservations from your profile.

### Admin Panel
- Admin users can add, update, and delete movies and showtimes.
- Use the drag-and-drop interface to manage showtimes easily.

## Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

