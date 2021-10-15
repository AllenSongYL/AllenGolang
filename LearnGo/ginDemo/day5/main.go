package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// get传值
	r.GET("/", func(ct *gin.Context) {
		username := ct.Query("username")
		age := ct.Query("age")
		page := ct.DefaultQuery("sex", "1")
		ct.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	r.Run(":9400")
}
