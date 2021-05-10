package task_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var baseUrl string = "http://127.0.0.1:8000"

func TestGetTask(t *testing.T) {
	// 创建请求地址
	url := baseUrl + "/v1/gettask"
	// 创建请求体
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	// 创建请求客户端
	client := new(http.Client)
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// 结束前将请求断开否则会引发内存泄漏
	defer resp.Body.Close()
	t.Log(fmt.Sprintf("Response status is : %s", resp.Status))
	// 读取返回体
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	t.Log("Response is solved as follow:")
	// 转换返回体为 map[string]
	var data map[string]interface{}
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		t.Log(fmt.Sprintf("{key: %s, value: %v}", k, v))
	}
}
