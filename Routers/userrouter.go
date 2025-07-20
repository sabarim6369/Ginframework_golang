package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarim6369/Ginframework_golang/controllers"  // replace "your-module-name" with your actual module name
)

func UserRoutes(r *gin.Engine) {
	r.POST("/create", controllers.CreateUser)
}