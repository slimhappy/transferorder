package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transferorder/v1/action"
)

type req struct {
	Msg string `json:"msg"`
	Task int `json:"task"`
}

type resp struct {
	Msg string `json:"msg"`
	Action int `json:"action"`
	ErrorCode action.ErrorCodeType `json:"errorCode"`
}

func GetTask(c *gin.Context)  {
	var res resp
	if len(action.TaskQueue) >0{
		res.Msg="Task To do"
		res.ErrorCode = action.NoError
		res.Action = action.TaskQueue[0].Action
		action.TaskQueue = action.TaskQueue[1:]
		c.JSON(http.StatusOK,res)
	}else{
		res.Msg="No Task To do"
		res.ErrorCode = action.NoError
		res.Action = 0
		c.JSON(http.StatusOK,res)
	}
}