package main

import (
	"log"
	"os"

	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	config := db.GenConfig()
	DB, err := db.NewConnection(config)

	if err != nil {
		log.Fatal("Error in Opening Connection")
	}

	db.AutoMigrateAll(DB)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.TestRoute(router)
	routes.ProblemRoutes(router)


	router.Run(":"+port)

}