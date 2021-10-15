package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func middlefunc(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-中间件")
	c.Next() // 不会先执行下面的fmt，而是跳转到下一个func()
	fmt.Println("2-中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {
	r := gin.Default()

	// get传值
	r.GET("/", func(ct *gin.Context) {
		// 中间件
		// 返回前，执行一些操作
		fmt.Println("aaa")
	}, func(c *gin.Context) {
		c.String(200, "gin首页")
	})

	r.GET("/middle", middlefunc, func(c *gin.Context) {
		c.String(200, "gin首页")
	})
	r.Run(":9400")
}
