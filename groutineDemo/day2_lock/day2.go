package main

import (
	"fmt"
	"sync"
	"time"
)

// 阶乘 4!=4x3x2x1=24
//      7!=7x6x5x4x3x2x1=5040

// 声明一个全局变量
var (
	results = make(map[uint64]uint64, 10)
	// 声明一个全局的互斥锁
	// lock是一个全局的互斥锁
	// sync 是包
	// Mutex 互斥
	lock sync.Mutex
)

func Jx(i uint64) {
	var res uint64 = 1
	var a uint64
	for a = 1; a <= i; a++ {
		res *= a
	}
	lock.Lock()
	results[i] = res

	lock.Unlock()
}

func main() {
	fmt.Println("======Running======")
	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)

	//

	// 计算1-200的阶乘，，把各个数的阶乘放入map
	// 最后输出结果，使用goroutine

	var i uint64
	for i = 1; i <= 30; i++ {
		go Jx(i)
	}

	//time.Sleep(time.Second * 10)
	fmt.Printf("results: %v\n", results)

	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
	fmt.Println("======Ending======")
}
