package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

type PutMessage struct {
	Ip      string `json:"ip"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()
	// gin.SetMode(gin.ReleaseMode)

	// 根路由，查看API服务状态
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	r.POST("/post/log", func(context *gin.Context) {

		var pm PutMessage

		// 设置默认结构体
		var defaultPutMessage PutMessage
		defaultPutMessage.Ip = "1.1.1.1"
		defaultPutMessage.Time = "20210101_130000"
		defaultPutMessage.Message = "someMessage"

		if err := context.BindJSON(&pm); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"ERR":    "bind json err",
				"likeme": &defaultPutMessage,
			})
		} else if net.ParseIP(pm.Ip) == nil {
			context.JSON(http.StatusOK, gin.H{
				"ERR":     "err ip format",
				"ip-like": defaultPutMessage.Ip,
			})
		} else if reg, _ := regexp.MatchString("[0-9]{8}_[0-9]{6}", pm.Time); reg != true {
			context.JSON(http.StatusOK, gin.H{
				"ERR":       "err time format",
				"time-like": defaultPutMessage.Time,
			})
		} else if strings.Contains(pm.Message, "cpuFreePer") &&
			strings.Contains(pm.Message, "swapInfo") &&
			strings.Contains(pm.Message, "ipaddrs") &&
			strings.Contains(pm.Message, "diskInfo") &&
			strings.Contains(pm.Message, "||") {

			// 创建对应的文件
			var rootDir = "G:\\GO\\笔记\\AllenGolang\\MyScripts\\XunJianAPI\\logtest"
			ipdir := path.Join(rootDir, pm.Ip)

			if _, err := os.Stat(ipdir); err != nil {
				os.MkdirAll(ipdir, 0644)
			}
			timefile := path.Join(ipdir, pm.Time+".log")
			tf, err := os.Create(timefile)
			if err != nil {
				fmt.Println("创建文件失败：", timefile)
			}
			defer tf.Close()
			if _, err := tf.WriteString(pm.Message); err == nil {
				// 返回正确post请求
				context.JSON(http.StatusOK, gin.H{
					"message":  "file create success",
					"filename": timefile,
				})
			}
		} else {
			context.JSON(http.StatusOK, gin.H{
				"ERR": "err message format",
			})

		}
	})
	r.Run(":8999")
}
