package main

import (
	"transferorder/config"
	"transferorder/router"
)

func main() {
	// 加载router
	r := router.Router
	// 读取服务端口号
	p := config.Conf.Svr.Port
	// 启动服务
	err := r.Run(":" + p)
	if err != nil {
		panic(err)
	}
}
