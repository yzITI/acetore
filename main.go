package main

import (
    "github.com/gin-gonic/gin"
    "acetore/controller"
    "net/http"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.POST("/", controller.Receive)
    router.StaticFS("/file", http.Dir("public"))
	_ = router.Run(":5000")
}
