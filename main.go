package main

import (
	"test-go-server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/parameter/valuebyobject/:main", controllers.GetValueByObject)
	r.GET("/api/v1/parameter/formatbyobject/:main", controllers.GetFormatByObject)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
