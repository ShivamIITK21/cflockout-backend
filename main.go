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
	routes.LockoutRoutes(router)


	router.Run(":"+port)

	// var z map[string]string = make(map[string]string)
	// z["Yuvraj"] = "23658"
	// z["Shivam"] = "63218"
	// z["Devansh"] = "23908"

	// var prob models.Problem
	// prob.ContestID = 3256
	// s := "A"
	// prob.Index = &s
	// prob.Name = &s
	// prob.Rating = 1678
	// prob.Tags = &s
	// var zz []models.ProblemInfo
	// var k models.ProblemInfo
	// k.Task = prob
	// k.Score = &s
	// k.FirstSolvedBy = &s
	// zz = append(zz, k)
	// k.Task = prob
	// k.Score = &s
	// k.FirstSolvedBy = &s
	// zz = append(zz, k)
	// k.Task = prob
	// k.Score = &s
	// k.FirstSolvedBy = &s
	// zz = append(zz, k)

	// var sd models.SessionData
	// sd.Participants = &z
	// sd.Problems = datatypes.NewJSONType(zz)

	// lockout := models.Lockout{
	// 	SessionId: &s,
	// 	SessionData: datatypes.NewJSONType(sd),
	// }
	// db.DB.Create(lockout)

	// var found models.Lockout

	// db.DB.First(&found)
	// fmt.Println(found)

}
