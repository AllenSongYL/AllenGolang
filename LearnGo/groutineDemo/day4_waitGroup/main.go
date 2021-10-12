package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func printInfo(s string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Info = ", s, "i= ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)
	wg.Add(1)
	go printInfo("Golang")
	wg.Wait()
}
