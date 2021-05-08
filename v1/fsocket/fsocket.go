package fsocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DownloadFile(c *gin.Context) {
	im:="/data/cache.png"
	c.File(im)
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("upload")
	dst := "/data/cache.png"
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
