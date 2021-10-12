package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`

	Password string `form:"password" json:"password"`
}

func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
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
	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "time.html", gin.H{
			"data": data,
		})
	})

	r.GET("/", func(context *gin.Context) {
		username := context.Query("username")
		age := context.Query("age")
		page := context.DefaultQuery("page", "1")

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

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
			"age":      age,
		})
	})

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

	r.Run(":9300")
}
