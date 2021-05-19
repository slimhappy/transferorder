package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 单个任务结构体
type Task struct {
	Action int
}

// 任务序列
var TaskQueue []Task

// 错误码类型
type ErrorCodeType int

// 返回体中的错误码变量
const (
	NoError         ErrorCodeType = 0
	ErrorInvalidReq ErrorCodeType = 1
)

// 请求体
type req struct {
	Msg    string `json:"msg"`
	Action int    `json:"action"`
}

// 返回体
type resp struct {
	Msg       string        `json:"msg"`
	Action    int           `json:"action"`
	ErrorCode ErrorCodeType `json:"errorCode"`
}

func handleGetTaskTaskToDo(p *resp) {
	if len(TaskQueue) > 0 {
		p.Msg = "Task To do"
		p.ErrorCode = NoError
		p.Action = TaskQueue[0].Action
		TaskQueue = TaskQueue[1:]
	}
}

func handleGetTaskNoTask(p *resp) {
	p.Msg = "No Task To do"
	p.ErrorCode = NoError
	p.Action = 0
}

// /v1/gettask 获取任务
func GetTask(c *gin.Context) {
	var p resp
	handleGetTaskNoTask(&p)
	handleGetTaskTaskToDo(&p)
	c.JSON(http.StatusOK, p)
}

// /v1/pushtask 推送任务
func PushTask(c *gin.Context) {
	var r req
	var p resp
	err := c.ShouldBindJSON(&r)
	if err != nil {
		p.ErrorCode = ErrorInvalidReq
		p.Msg = "Invalid Request"
		c.JSON(http.StatusBadRequest, p)
	} else {
		p.ErrorCode = NoError
		p.Msg = "Request Accept"
		p.Action = r.Action
		c.JSON(http.StatusOK, p)
		TaskQueue = append(TaskQueue, Task{Action: r.Action})
	}
}
