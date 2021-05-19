package fsocket

import (
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	im := "/data/cache.png"
	c.File(im)
}
