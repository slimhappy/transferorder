package fsocket_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

var baseUrl string = "http://127.0.0.1:8000"
var filePath string = "/Users/yizhang/Downloads/test.png"

func TestUploadFile(t *testing.T) {
	// 创建请求地址
	url := baseUrl + "/v1/upload"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("upload", filePath)
	if err != nil {
		fmt.Println("error writing to buffer")
		panic(err)
	}
	// iocopy
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		panic(err)
	}
	// 结束前将请求断开否则会引发内存泄漏
	defer resp.Body.Close()
	t.Log(fmt.Sprintf("Response status is : %s", resp.Status))
	// 读取返回体
	jsonStr, err := ioutil.ReadAll(resp.Body)
	t.Log(fmt.Sprintf("%s", jsonStr))
}
