package fsocket

import (
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	im:="./imcache/cache.png"
	c.File(im)
}
