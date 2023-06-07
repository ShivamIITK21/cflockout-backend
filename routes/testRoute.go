package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/ShivamIITK21/cflockout-backend/controllers"
)

func TestRoute(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/test", controllers.Test())
}