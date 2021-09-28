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



```
r.GET("", func(context *gin.Context) {
   context.String(200, "Hello World!")
})
```



## 热加载

使用第三方模块

模块一：go get -u github.com/pilu/fresh

下载完成后，进到文件目录下，执行 fresh



![image-20210927144920768](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20210927144920768.png)

![image-20210927145150730](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20210927145150730.png)



模块二：go get -u github.com/codegangsta/gin



## 返回值

- c.String()

- c.JSON()

- c.JSONP()

  在url中使用?callback=xxx时，会把xxx放在json格式前面

  执行xxx方法，传入json数据。

  主要用来解决跨域问题

- c.XML()

- c.HTML()



## 路由



使用 GET, POST, PUT, PATCH, DELETE 和OPTIONS

```
func main() {
	// 使用默认中间件创建GIN路由器:
	// 记录器和恢复（无崩溃）中间件
	router := gin.Default()
	
	// GET请求：获取数据
	router.GET("/someGet", getting)
	
	// POST请求：主要用于增加数据
	router.POST("/somePost", posting)
	
	// PUT请求：主要用于编辑数据
	router.PUT("/somePut", putting)
	
	// DELETE请求：主要用于删除数据
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

