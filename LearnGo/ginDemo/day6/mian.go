package main

import (
	"github.com/gin-gonic/gin"
	"go_env/LearnGo/ginDemo/day6/controllers/api"
	"net/http"
)

func main() {
	r := gin.Default()

	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})
	}

	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/add", api.ApiController{}.Add)
		apiRouter.GET("/del", api.ApiController{}.Del)
		apiRouter.GET("/edit", api.ApiController{}.Edit)
	}

	r.Run(":9400")
}
