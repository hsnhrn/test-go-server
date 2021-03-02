package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", controllers.test)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
