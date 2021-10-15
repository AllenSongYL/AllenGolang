package baseControl

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (con BaseController) Success(c *gin.Context) {
	c.String(200, "Sueecss")
}

func (con BaseController) Error(c *gin.Context) {
	c.String(404, "Error")
}
