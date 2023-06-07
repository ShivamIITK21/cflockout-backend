package main

import (
	"log"
	"os"

	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/models"
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

	i1 := "A"
	i2 := "B"
	n1 := "edasuilb"
	z := "dsj"

	problem := &models.Problem{
		ContestID: 1800,
		Index: &i1,
		Name: &n1,
		Ptype: &z,
	}
	problem2 := &models.Problem{
		ContestID: 1800,
		Index: &i2,
		Name: &n1,
		Ptype: &z,
	}

	err = DB.Create(problem).Error

	if err != nil {
		log.Fatal("MAA KI CHUU")
	}

	err = DB.Create(problem2).Error

	if err != nil {
		log.Fatal("MAA KI CHUU")
	}
	// DB.Create(problem)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.TestRoute(router)


	router.Run(":"+port)

}