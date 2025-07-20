package main

import (
	
	"github.com/gin-gonic/gin"
	"github.com/sabarim6369/Ginframework_golang/routers"
	"github.com/sabarim6369/Ginframework_golang/db"
)


func main() {
	r := gin.Default()
	db.ConnectToDB()
	routers.UserRoutes(r)



	r.Run()
}
