package main

import (
	"fmt"
	"go_env/TheAIgorithms/insertionSort"
	"math/rand"
	"time"
)

func main() {
	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)

	var Aslice []int = make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		Aslice[i] = rand.Int()
	}
	fmt.Println(Aslice)
	a := insertionSort.InsertionSort(Aslice)
	fmt.Println(a)

	// 程序结束输出
	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
}
