package main

import (
	
	"github.com/gin-gonic/gin"
	"github.com/sabarim6369/Ginframework_golang/routers"
)


func main() {
	r := gin.Default()
	routers.UserRoutes(r)



	r.Run()
}
