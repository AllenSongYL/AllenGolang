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

## 返回HTML模板
html模板文件中可以使用 `{{ .变量名 }}` 输出值

### **定义变量**

{{ $t := .title}}

### 比较
eq ==
ne !=
lt <
le <=
gt >
ge >=

### 条件判断
{{ if xx}} T1 {{ else if xxx}} T0 {{end}}



### 自定义模板函数

在配置引擎的下边

r.SetFuncMap(template.FuncMap{

})

main.go

~~~
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp,0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	fmt.Println("success")
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,

	})
	r.LoadHTMLGlob("templates/*")
	data := time.Now().Unix()

	fmt.Println(data)
	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "time.html", gin.H{
			"data": data,
		})
	})
	r.Run(":9300")
}
~~~

time.html

~~~
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>当前时间</title>
</head>
<body>
 <h2>{{  .data }}</h2>
  <h2>{{ UnixToTime .data }}</h2>
</body>
</html>
~~~

![image-20211006071419538](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20211006071419538.png)



## GET POST传值

### GET传值

~~~
r.GET("/", func(context *gin.Context) {
		username := context.Query("username")
		age := context.Query("age")
		page := context.DefaultQuery("page", "1")

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"age": age,
			"page": page,
		})
	})
~~~

http://127.0.0.1:9300/?username=zhangsan&age=10

![image-20211006072906686](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20211006072906686.png)

![image-20211006072932782](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20211006072932782.png)

### POST传值

```
// POST
r.GET("/user", func(c *gin.Context) {
   c.HTML(http.StatusOK, "user.html", gin.H{})
})
r.POST("/doAddUser", func(c *gin.Context) {
   // html模板中的username
   username := c.PostForm("username")
   password := c.PostForm("password")
   age := c.DefaultQuery("age", "18")
   c.JSON(http.StatusOK, gin.H{
      "username": username,
      "password": password,
      "age": age,
   })
})
```

user.html

```
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>当前时间</title>
</head>
<body>

    <form action="/doAddUser", method="post">
        用户名:<input type="text" name="username"/> <br>
        密码:<input type="password" name="password"/> <br>
        <input type="submit" value="提交">
    </form>
</body>
</html>
```



### 传值绑定到结构体

.ShouldBind() 能够基于请求自动提取JSON,from表单和QueryString类型的数据，并把值绑定到指定的结构体对象



```
// 定义结构体
type UserInfo struct {

   Username string `form:"username" json:"username"`

   Password string `form:"password" json:"password"`

}
```

~~~

r.GET("/getuser", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
	})
	
~~~

### 获取POST XML数据

.getRawData         // 从c.Request.Bodu 读取请求数据

![getRawData](G:\GO\笔记\AllenGolang\go笔记\gin笔记\图\getRawData.png)
