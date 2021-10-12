package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Artic struct {
	Title   string
	Message string
	Name    string
}

func main() {
	r := gin.Default()
	// 加载所有模板文件 必须放在引擎的下面
	r.LoadHTMLGlob("templates/*")

	// 这个需要配置每个模板，不常用
	//r.LoadHTMLFiles("templates/new.html", "templates/goods.html")

	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "new.html", gin.H{
			"title": "我是后台模板文件news.html",
		})
	})

	r.GET("/goods", func(context *gin.Context) {
		context.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品页面",
			"price": 20,
			"name":  "纸巾",
		})
	})

	// 结构体
	r.GET("/structs", func(context *gin.Context) {
		Astructs := &Artic{
			Title:   "aaa",
			Message: "bbb",
			Name:    "ccc",
		}
		context.HTML(http.StatusOK, "structs.html", gin.H{
			"all": Astructs,
		})
	})

	r.Run(":12302")
}
