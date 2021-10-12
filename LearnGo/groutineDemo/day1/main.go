package main

import (
	"fmt"
	"strconv"
	"time"
)

func HelloWorld() {
	for i := 1; i < 11; i++ {
		fmt.Println("HelloWorld():Hello,Golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)

	// runtime 包的NumCPU() 获取cpu数
	//num := runtime.NumCPU()
	//fmt.Println("CPU数：", num)
	// 设置最大数量
	//runtime.GOMAXPROCS(num -1)

	go HelloWorld()

	for i := 1; i < 11; i++ {
		fmt.Println("mian(): Hello,Golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	// 程序结束输出
	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
}
