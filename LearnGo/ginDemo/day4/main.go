package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ct *gin.Context) {
		ct.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
		})
	})
	r.Run(":9400")
}
