package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env") //Access environment variables
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	serverPort := os.Getenv("SERVER_PORT") //Access server port

	router := mux.NewRouter() //Create a new router

	log.Println("Server is running")
	err = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), router) //This will run the server and listen for requests
	if err != nil { 
		log.Fatalf("Error occured in server: %s", err)
	}

}