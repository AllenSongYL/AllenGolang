package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("开始时间：", timeStartFormat)

	// 创建一个默认的路由引擎
	r := gin.Default()

	// /ping 路由，返回匿名函数，
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 主路由，返回"Hello World!"
	r.GET("", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	//
	r.GET("/pong", func(context *gin.Context) {
		// gin.H{} = type H map[string]interface{}
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "ping",
		})
	})

	// r.Run()    不指定地址运行，默认端口8080
	r.Run("localhost:12302")
}
