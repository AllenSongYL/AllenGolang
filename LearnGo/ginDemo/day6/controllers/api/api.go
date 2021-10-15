package api

import (
	"github.com/gin-gonic/gin"
	"go_env/LearnGo/ginDemo/day6/controllers/baseControl"
)

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

// 函数方式
//func Apidel(ct *gin.Context) {
//	ct.String(http.StatusOK, "api/del")
//}
//
//func Apiadd(ct *gin.Context) {
//	ct.String(http.StatusOK, "api/add")
//}
//
//func Apiedit(ct *gin.Context) {
//	ct.String(http.StatusOK, "api/edit")
//}
