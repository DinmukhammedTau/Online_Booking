package main

import (
	"log"
	"net/http"
	"onlinebooking/handlers"
	"onlinebooking/models"

	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connectDB() *gorm.DB {
	// Connect to the database
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Run database migrations
	database.AutoMigrate(&models.Booking{})
	log.Println("Database connected and migrated")
	return database
}

func main() {
	// Connect to the database
	db := connectDB()
	handlers.InitDB(db) // Pass database instance to handlers

	// Set up the router
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// CRUD routes
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
