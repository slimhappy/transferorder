package router

import (
	"github.com/gin-gonic/gin"
	"transferorder/v1/action"
	"transferorder/v1/fsocket"
	"transferorder/v1/task"
)
var Router = gin.New()

func init() {
	v1 := Router.Group("/v1")
	{
		v1.POST("/upload", fsocket.UploadFile)
		v1.GET("/download",fsocket.DownloadFile)
		v1.POST("/action",action.Action)
		v1.GET("/gettask",task.GetTask)
	}
}