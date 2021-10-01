package main

import (
	"fmt"
	"sync"
)

func A(i int) {
	fmt.Println("我是A", i)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("我是main")
	wg.Add(1)
	for a := 0; a < 10; a++ {
		go A(a)
	}

	wg.Wait()
	fmt.Println("执行完了")

}
