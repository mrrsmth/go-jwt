package main

import (
	_ "go-jwt/controllers"
	"go-jwt/database"
	_ "go-jwt/middlewares"
	"go-jwt/models"
	"log"

	_ "github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	err := database.InitDatabase()
	if err != nil {
		// Log the error and exit
		log.Fatalln("could not create database", err)
	}
	// Automigrate the User model
	// AutoMigrate() automatically migrates our schema, to keep our schema upto date.
	database.GlobalDB.AutoMigrate(&models.User{})
}
