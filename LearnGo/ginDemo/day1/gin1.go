package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义结构体
type headTest struct {
	// `` 添加标签 显示结构体属性的样式
	Title string `json:"title"`
	// 字母开头小写，导致无法访问到
	Message string `json:"message"`
	Name    string `json:"name"`
}

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

	// 使用map返回json
	r.GET("/json1", func(context *gin.Context) {
		// gin.H{} = type H map[string]interface{}
		context.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "map[string]interface{}",
		})
	})

	// 测似gin.H返回json
	r.GET("/json2", func(context *gin.Context) {
		// gin.H{} = type H map[string]interface{}
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "gin.H",
		})
	})

	// 测似通过结构体返回json
	r.GET("/json3", func(context *gin.Context) {
		structA := &headTest{
			Title:   "首页",
			Message: "这是一个结构体测试",
			Name:    "allen",
		}
		context.JSON(http.StatusOK, structA)
	})

	// 响应JSONP 可以传入回调函数 主要用来解决跨域问题
	// 当url为 http://127.0.0.1:12302/jsonp?callback=xxx
	// 返回：xxx( {"title": "jsonp首页","message": "这是一个jsonp","name": "allen"})
	r.GET("/jsonp", func(context *gin.Context) {
		structA := &headTest{
			Title:   "jsonp首页",
			Message: "这是一个jsonp",
			Name:    "allen",
		}
		context.JSONP(http.StatusOK, structA)
	})

	// 返回XML
	r.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"success": true,
			"message": "i`m XML",
		})
	})

	// r.Run()    不指定地址运行，默认端口8080
	r.Run("localhost:12302")
}
