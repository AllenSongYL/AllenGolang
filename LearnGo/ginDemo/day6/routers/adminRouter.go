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
