package fsocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("upload")
	log.Println(file.Filename)
	dst := "./imcache/cache.png"
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
