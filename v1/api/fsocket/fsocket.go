package fsocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type uploadResp struct {
	Msg       string `json:"msg"`
	ErrorCode int    `json:"action"`
}

func DownloadFile(c *gin.Context) {
	im := "/data/cache.png"
	c.File(im)
}

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("upload")
	if err != nil {

	}
	dst := "./cache.png"
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
