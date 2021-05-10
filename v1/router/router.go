package router

import (
	"github.com/gin-gonic/gin"
	"transferorder/v1/fsocket"
	"transferorder/v1/task"
)

var Router = gin.Default()

func init() {
	v1 := Router.Group("/v1")
	{
		// 图片上传接口
		v1.POST("/upload", fsocket.UploadFile)
		// 图片下载接口
		v1.GET("/download", fsocket.DownloadFile)
		// 上传任务接口
		v1.POST("/pushtask", task.PushTask)
		// 拉取任务接口
		v1.GET("/gettask", task.GetTask)
	}
}
