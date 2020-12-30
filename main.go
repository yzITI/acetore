package main

import (
	"acetore/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/", controller.Receive)
	router.GET("/verify/:token", controller.Verify)
	router.StaticFS("/file", http.Dir("public"))
	_ = router.Run(":5000")
}
