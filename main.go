package main

import (
	"os"

	"github.com/ShivamIITK21/cflockout-backend/routes"
	"github.com/gin-gonic/gin"
	
)


func main(){
	


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