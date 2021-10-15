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



## HTML模板渲染



```Go
r.LoadHTMLGlob("templates/*")
```



### 输出数据

{{ .变量名 }}



### 模板使用

~~~html
// 定义变量
{{ $t := .title }}

// 条件判断
{{ if ge .score 60 }}
	<p>及格</p>
{{else}}
	<p>及格</p>
{{end}}

// 循环遍历
{{ range $key,$value := .hobby }}
	<li>{{$key}}---{{value}}</li>
{{else}}
	<li>数组中没有数据</li>
{{end}}
~~~



### 模板放不同目录

通过define定义名称

~~~
{{define "admin/index.html"}}


{{end}}
~~~







### with解构结构体

~~~
{{.news.Content}}
{{.news.Title}}

{{ with .news}}
	{{.Content}}
	{{.Title}}
{{ end }}
~~~



### 预定义函数

len  		 {{ len .变量名 }}

and  		and x y 等价于 if x then y else x

or			 or x y    等价于 if x then x else y

not           返回他的单个参数的布尔值的否定

index       索引/键值



### 自定义函数



```go
func UnixToTime(timestamp int64) string {
   t := time.Unix(timestamp, 0)
   return t.Format("2006-01-02 15:04:05")
}

// 自定义模板函数 放在加载模板前
r.SetFuncMap(template.FuncMap{
    
		"UnixToTime": UnixToTime,
	})

data := time.Now().Unix()
	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "time.html", gin.H{
			"data": data,
		})
	})
```

~~~html
{{UnixToTime }}
~~~



### 外部引入

{{ template "xxx/xxx"  . }}



### 静态WEB服务

前面的static路由，后面的static为路径

~~~
r.Static("/static", "./static")
~~~







### 返回HTML模板

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

**传入结构体地址**

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

.getRawData         // 从c.Request.Bodu 读取请求数据。返回byte类型的切片

![getRawData](G:\GO\笔记\AllenGolang\go笔记\gin笔记\图\getRawData.png)



# 路由组



### 简单的路由组

~~~go
v1 := router.Group("/v1")
~~~

~~~
apiRouter := r.Group("/api")
	{
		apiRouter.GET("/add", func(ct *gin.Context) {
			ct.String(http.StatusOK, "add")
		})
		apiRouter.GET("/del", func(ct *gin.Context) {
			ct.String(http.StatusOK, "del")
		})
		apiRouter.GET("/edit", func(ct *gin.Context) {
			ct.String(http.StatusOK, "edit")
		})
	}
~~~



### 文件形式

main.go

~~~go
// 导入routers包
r := gin.Default()
routers.AdminRouter(r)
~~~



创建一个routers文件包

apiRouters.go

adminRouters.go

~~~go
package routers

import "github.com/gin-gonic/gin"

func AdminRouter(r *gin.Engine) {
	adminRouter := r.Group("admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.String(200, "/admin")
		})
		adminRouter.GET("/edit", func(c *gin.Context) {
			c.String(200, "/admin/edit")
		})
		adminRouter.GET("/add", func(c *gin.Context) {
			c.String(200, "/admin/add")
		})
		adminRouter.GET("/list", func(c *gin.Context) {
			c.String(200, "/admin/list")
		})
	}
}
~~~

defaultRouters.go

类似adminRouters.go



## 自定义控制器

### 通过函数

main.go

~~~go
apiRouter := r.Group("/api")
	{
		apiRouter.GET("/add", api.Apiadd)
		apiRouter.GET("/del", api.Apidel)
		apiRouter.GET("/edit", api.Apiedit)
	}
~~~

新建controllers文件夹，在新建api子文件夹，最后创建api.go

~~~go
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Apidel(ct *gin.Context) {
	ct.String(http.StatusOK, "api/del")
}

func Apiadd(ct *gin.Context) {
	ct.String(http.StatusOK, "api/add")
}

func Apiedit(ct *gin.Context) {
	ct.String(http.StatusOK, "api/edit")
}
~~~



### 通过结构体

main.go

~~~
apiRouter := r.Group("/api")
	{
		apiRouter.GET("/add", api.ApiController{}.Add)
		apiRouter.GET("/del", api.ApiController{}.Del)
		apiRouter.GET("/edit", api.ApiController{}.Edit)
	}
~~~

api.go

~~~
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

func (con ApiController) Del(ct *gin.Context) {
	ct.String(http.StatusOK, "api/del")
}

func (con ApiController) Add(ct *gin.Context) {
	ct.String(http.StatusOK, "api/add")
}

func (con ApiController) Edit(ct *gin.Context) {
	ct.String(http.StatusOK, "api/edit")
}
~~~



### 控制器继承

新建BaseController

~~~go
package baseControl

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (con BaseController) Success(c *gin.Context){
	c.String(200, "Sueecss")
}

func (con BaseController) Error(c *gin.Context){
	c.String(404, "Error")
}
~~~



NewCOntroller 继承 BaseController

继承后就可以调用控制器里面的公共方法了

~~~
// 结构体方法
type ApiController struct {
	baseControl.BaseController
}

func (con ApiController) Del(ct *gin.Context) {
	con.Success(ct)
}

func (con ApiController) Add(ct *gin.Context) {
	con.Success(ct)
}

func (con ApiController) Edit(ct *gin.Context) {
	con.Success(ct)
}
~~~



## 中间件

r := gin.Default()

默认加载Logger和Recovery两个中间件

- Logger中间件
- 

不想使用这两个直接使用New

~~~go
r.GET("/", func(ct *gin.Context) {
		// 中间件
		// 回复请求前，执行一些操作
		fmt.Println("aaa")
	}, func(c *gin.Context) {
		c.String(200, "gin首页")
	})

r.GET("/", middlefunc, func(c *gin.Context) {
		c.String(200, "gin首页")
	})
~~~





### context.Next()    

// 调用该请求剩余处理程序



~~~go
func middlefunc(c *gin.Context) {
    start := time.Now().UnixNano()
	fmt.Println("1-中间件")
    c.Next()   // 不会先执行下面的fmt，而是跳转到下一个func()
	fmt.Println("2-中间件")
    end := time.Now().UnixNano()
    fmt.Println(end.Sub(start))
}

// 输出
//17:10:11 app         | 1-中间件
//17:10:11 app         | 2-中间件
//513700

~~~



### context.Abort()

// 终止调用该请求的剩余处理程序

会执行完该中间件，不会执行后面的程序



### 全局中间件

r.Use()

~~~
r.Use(middleOne, middleTwo)
~~~

### 路由分组中配置中间件

- r.Group("/api", xxx)   中添加中间件
- adminRouters.Use(xxx)  添加

### 中间件和控制器之间共享数据

ctx.Set("username","zs")

username,_ := ctx.Get("username")
