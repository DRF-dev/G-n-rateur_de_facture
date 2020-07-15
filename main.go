package main

import (
	"downloadFile/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	api := r.Group("/api")
	{
		api.POST("/add", controllers.Add)
	}
	r.GET("/output/:file")
	r.Run(":4000")
}
