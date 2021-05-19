package fsocket

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"regexp"
	"time"
	"transferorder/config"
	"transferorder/v1/common"
)

// /v1/upload handler
func UploadPic(c *gin.Context) {
	var (
		err  error
		resp = &common.Response{}
	)

	defer func() {
		common.HandleResponse(c, err, resp)
	}()

	file, err := c.FormFile("file")
	if err != nil {
		return
	}
	ct := file.Header.Get("Content-Type")
	if ct == "" {
		resp.SetCode(-1)
		resp.SetMsg("Can't get file's Content-Type")
		return
	}
	mf, err := regexp.MatchString(config.Conf.Svr.FileTypePattern, ct)
	if err != nil {
		return
	}
	if !mf {
		resp.SetCode(-1)
		resp.SetMsg("File is not a pic!")
		return
	}
	fd := config.Conf.Svr.FileDir
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(fd, folderName)
	savePath := folderPath + "/" + filepath.Base(file.Filename)
	err = handleUploadFileMkdir(folderPath)
	if err != nil {
		return
	}
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		return
	}
}

// 处理创建文件夹
func handleUploadFileMkdir(fp string) (err error) {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		err = os.Mkdir(fp, os.ModePerm)
		if err != nil {
			return err
		}
		err = os.Chmod(fp, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
