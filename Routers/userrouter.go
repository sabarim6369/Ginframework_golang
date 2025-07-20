package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarim6369/Ginframework_golang/controllers"  
)

func UserRoutes(r *gin.Engine) {
	r.POST("/create", controllers.CreateUser)
}