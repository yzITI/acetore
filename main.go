package main

import (
	"acetore/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/upload/:hash", controller.VerifyUpload)
	router.POST("/upload", controller.Upload)
	router.StaticFS("/file", http.Dir("public"))
	_ = router.Run(":5000")
}
