package main

import (
	"acetore/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/upload/:hash", controller.verifyUpload)
	router.POST("/upload", controller.Upload)
	router.StaticFS("/file", http.Dir("public"))
	_ = router.Run(":5000")
}
