package Mytime

import (
	"fmt"
	"time"
)

func StartTime() time.Time {
	timeStart := time.Now()
	fmt.Println("程序开始运行......")
	fmt.Println("开始时间： ", timeStart.Format("2006-01-02 15:04:05"))
	return timeStart
}

func EndTime() time.Time {
	timeEnd := time.Now()
	fmt.Println("程序运行结束。")
	fmt.Println("结束时间： ", timeEnd.Format("2006-01-02 15:04:05"))
	return timeEnd
}
