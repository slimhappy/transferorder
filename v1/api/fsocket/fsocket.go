package fsocket

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"transferorder/config"
)

type uploadResp struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func DownloadFile(c *gin.Context) {
	im := "/data/cache.png"
	c.File(im)
}

func handleUploadFileError(p *uploadResp, err error) (f bool) {
	if err != nil {
		p.Msg = err.Error()
		p.Code = 1
		return false
	}
	return true
}

// 处理正常的返回
func handleUploadFileOk(p *uploadResp) {
	p.Msg = "upload ok"
	p.Code = 0
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

func UploadFile(c *gin.Context) {
	var p uploadResp
	for ok := true; ok; ok = false {
		// 使用gin库读取参数中的文件
		handleUploadFileOk(&p)
		file, err := c.FormFile("upload")
		if !handleUploadFileError(&p, err) {
			break
		}
		// 创建按日期分开的文件夹
		fd := config.Conf.Svr.FileDir
		folderName := time.Now().Format("20060102")
		folderPath := filepath.Join(fd, folderName)
		savePath := folderPath + "/" + filepath.Base(file.Filename)
		err = handleUploadFileMkdir(folderPath)
		if !handleUploadFileError(&p, err) {
			break
		}
		// 将文件保存到对应文件夹
		err = c.SaveUploadedFile(file, savePath)
		if !handleUploadFileError(&p, err) {
			break
		}
	}
	c.JSON(http.StatusOK, p)
}
