package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Task struct{
	Action int
}

var TaskQueue []Task

type ErrorCodeType int
const (
	NoError         ErrorCodeType = 0
	ErrorInvalidReq ErrorCodeType = 1
)

type req struct {
	Msg string `json:"msg"`
	Action int `json:"action"`
}

type resp struct {
	Msg       string        `json:"msg"`
	ErrorCode ErrorCodeType `json:"errorCode"`
}

func Action(c *gin.Context) {
	var r req
	var res resp
	err := c.ShouldBindJSON(&r)
	if err != nil{
		res.ErrorCode = ErrorInvalidReq
		res.Msg = "Invalid Request"
		c.JSON(http.StatusBadRequest, res)
	}else {
		res.ErrorCode = NoError
		res.Msg = "Request Accept"
		c.JSON(http.StatusOK, res)
		TaskQueue = append(TaskQueue, Task{Action: r.Action})
	}
}