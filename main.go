package main

import (
	"os"

	"github.com/ShivamIITK21/cflockout-backend/routes"
	// "fmt"
	// "log"

	// "github.com/ShivamIITK21/cflockout-backend/db"
	// "github.com/ShivamIITK21/cflockout-backend/helpers"
	// "github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
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


	router.Run(":"+port)

	// var user models.User
	// z := "Rounak Sharma"
	// user.CFid = &z
	// user.Password = &z
	// user.Username = &z
	// user.UserType = &z
	// user.Token = &z

	// db.DB.Create(&user)
	
	// var x helpers.SignedDetails
	// x.Username = "Shivam"
	// x.CFid = "phoenix2913"
	// x.UserType = "deuiedsba"
	// token, _, err := helpers.GenerateTokens(x.Username, x.CFid, x.UserType)
	// if err!=nil {
	// 	log.Fatal("chud gaya")
	// }
	// helpers.UpdateAllTokens(token, z)
	// claims, _ := helpers.ValidateToken(token)
	// fmt.Println(claims)

}
