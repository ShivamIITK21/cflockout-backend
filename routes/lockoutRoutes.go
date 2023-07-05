package routes

import (
	"github.com/ShivamIITK21/cflockout-backend/controllers"
	"github.com/ShivamIITK21/cflockout-backend/middleware"
	"github.com/gin-gonic/gin"
)


func LockoutRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/lockout/create", middleware.Authorize(), controllers.CreateLockoutController())
	incomingRoutes.GET("/lockout", middleware.Authorize(), controllers.LockoutController())
	incomingRoutes.GET("/lockout/getUserRating", middleware.Authorize(), controllers.GetRating())
}