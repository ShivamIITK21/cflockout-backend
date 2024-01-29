package routes

import (
	"github.com/ShivamIITK21/cflockout-backend/controllers"
	"github.com/ShivamIITK21/cflockout-backend/middleware"
	"github.com/gin-gonic/gin"
)

func ProblemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/problems/refresh", middleware.Authorize(), controllers.RefreshController())
	incomingRoutes.GET("/problems", controllers.GetUserSolvedProblems())
}
