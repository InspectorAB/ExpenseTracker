package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB
var c *cors.Cors

func initDB() {
    var err error
    connStr := "user=postgres password=06124300 dbname=expense_tracker sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Successfully connected to the database!")
}

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}

func registerUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    sqlStatement := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
    err := db.QueryRow(sqlStatement, user.Username, user.Email, user.Password).Scan(&user.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.CreatedAt = time.Now()
    json.NewEncoder(w).Encode(user)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    var dbUser User
    sqlStatement := `SELECT id, username, email, password, created_at FROM users WHERE email=$1`
    row := db.QueryRow(sqlStatement, user.Email)
    err := row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Password, &dbUser.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    if dbUser.Password != user.Password {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(dbUser)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // This will display text in the browser when visiting localhost:8000
    fmt.Fprintf(w, "Welcome to the Expense Tracker API!")
}

func main() {
    initDB()
    defer db.Close()

    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/register", registerUser).Methods("POST")
    r.HandleFunc("/login", loginUser).Methods("POST")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"}, // React app URL
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type"},
    })
    http.Handle("/", c.Handler(r))
    fmt.Println("Server started ")
    log.Fatal(http.ListenAndServe(":8080", nil))

}
