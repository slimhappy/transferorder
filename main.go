package main

import (
	"github.com/gin-gonic/gin"
	"transferorder/v1/router"
)


func main() {
	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	r := router.Router
	err := r.Run(":8090")
	if err != nil {
		panic(err)
	}
}