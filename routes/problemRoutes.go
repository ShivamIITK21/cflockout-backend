package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/ShivamIITK21/cflockout-backend/controllers"
)

func ProblemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/problems/refresh", controllers.RefreshController())
}