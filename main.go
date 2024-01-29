package main

import (
	"os"
	// "fmt"
	
	// "github.com/ShivamIITK21/cflockout-backend/db"
	// "github.com/ShivamIITK21/cflockout-backend/helpers"
	// "gorm.io/datatypes"
	// "github.com/ShivamIITK21/cflockout-backend/helpers"
	// "github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/ShivamIITK21/cflockout-backend/routes"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func main(){

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
    router.Use(CORSMiddleware())

	routes.TestRoute(router)
	routes.ProblemRoutes(router)
	routes.AuthRoutes(router)
	routes.LockoutRoutes(router)


	router.Run(":"+port)

}
