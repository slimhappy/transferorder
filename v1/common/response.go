package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	// 0为正确，非0为发生错误
	Code int `json:"code"`
	// 发生错误时的错误消息
	Msg string `json:"msg"`
}

func (r *Response) SetCode(i int) {
	r.Code = i
}

func (r *Response) SetMsg(s string) {
	r.Msg = s
}

type IResponse interface {
	SetCode(int)
	SetMsg(string)
}

func HandleResponse(c *gin.Context, err error, resp IResponse) {
	if err != nil {
		resp.SetCode(-1)
		resp.SetMsg(err.Error())
		c.JSON(http.StatusAccepted, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
