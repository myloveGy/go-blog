package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog/app/routers"
	_ "blog/configs"
	"blog/global"
)

// @title GO博客项目
// @version 1.0
// @description GO编程之旅：一起用GO做项目
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
