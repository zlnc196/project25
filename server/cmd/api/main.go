package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"project25/internal/database"
)

func main() {
	err := godotenv.Load(".env") //Access environment variables
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	//Access database information from environment
	DATABSE_HOST := os.Getenv("DATABSE_HOST")
	DATABASE_PORT := os.Getenv("DATABASE_PORT")
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DATABSE_HOST, DATABASE_PORT, DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{}) //Accesses postgresql database
	if err != nil {
		log.Fatalf("failed connection to database due to: %s", err)
	}

	db.AutoMigrate(&database.User{}) //Automatically migrate Users
	
	SERVER_PORT := os.Getenv("SERVER_PORT") //Access server port

	router := mux.NewRouter() //Create a new router

	log.Println("Server is running")
	err = http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), router) //This will run the server and listen for requests
	if err != nil {
		log.Fatalf("Error occured in server: %s", err)
	}

}
