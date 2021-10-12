package main

import (
	"fmt"
	"go.uber.org/zap"
)

func main() {
	// 默认情况下只输出时间
	// SetFlags函数可以用来配置输出
	// Ldata 日期 Ltime 时间 Lmicroseconds微秒级别的时间 LUTC 使用UTC时间
	// Llongfile 文件全路径名+行号
	// Lshortfile 文件名+行号
	// LstdFlags = Ldate | Ltime  标准logger初始值
	//log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	//log.Println("这是一条很普通的日志!")
	//
	//// SetPrefix设置前缀
	//log.SetPrefix("[go]: ")
	//log.Println("这是一条很普通的日志!")
	//
	//// 学习zap日志库
	//fmt.Println("学习zap日志库")
	// Sugared Logger 支持结构化和printf风格的日志记录
	// Logger 只支持强类型的结构化日志记录 速度更快

	// 创建logger的方法
	// zap.NewExample()、zap.NewDevelopment()、zap.NewProduction()、zap.New()
	//fmt.Println("使用zap.NewExample()---")
	// 输出json格式
	//logger := zap.NewExample()
	//
	//// 将缓存同步到文件中
	//defer logger.Sync()
	//
	//url := "http://example.org/api"
	//logger.Info(
	//	"fail to fetch URL",
	//	zap.String("url",url),
	//	zap.Int("attempt",3),
	//	)
	//
	//logger.Info("xxfadf")
	//logger.Error("error")

	fmt.Println("使用zap.NewProduction---")
	// 输出json格式
	// DEBUG消息不记录
	// Error,Dpanic级别的记录，会在堆栈中跟踪文件，warn不会
	loggerDev, _ := zap.NewProduction()
	loggerDev.Info("infosfdsfag")
	loggerDev.Debug("debug111")
	loggerDev.Error("error!")
	loggerDev.Warn("warning!")
	loggerDev.Fatal("fatal")

}
