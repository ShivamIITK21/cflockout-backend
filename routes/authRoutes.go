package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/ShivamIITK21/cflockout-backend/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/auth/login", controllers.Login())
	incomingRoutes.POST("/auth/signup", controllers.Signup())
	incomingRoutes.GET("/auth/findUser", controllers.FindUser())
	incomingRoutes.GET("/auth/verifyToken", controllers.VerifyToken())
}

