# Gin

## 部署



```
go get -u github.com/gin-gonic/gin

```

简单例子

```
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//  使用常量，需要http.StatusOK
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```



## 路由



使用 GET, POST, PUT, PATCH, DELETE 和OPTIONS

```
func main() {
	// 使用默认中间件创建GIN路由器:
	// 记录器和恢复（无崩溃）中间件
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
```



## API参数

可以通过context的Param方法来获取API参数

```
r.GET("/user/:name/*action", func(c *gin.Context) {
name := c.Param("name")
action := c.Param("action")
message := name + " is " + action
c.String(http.StatusOK, message)
})
```
