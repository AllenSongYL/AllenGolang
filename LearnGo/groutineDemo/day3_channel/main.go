package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
	Age  int
}

func main() {

	// 使用管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println("intChan: ", intChan)
	fmt.Println(reflect.TypeOf(intChan))
	fmt.Printf("intChan的值: %v, intChan本身的地址：%p\n", intChan, &intChan)

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	close(intChan)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	// 读取数据
	var num2 int

	num2 = <-intChan
	fmt.Println(num2)
	fmt.Println(len(intChan))

	// 创建一个cat结构体管道
	catchan := make(chan cat, 10)
	catA := cat{
		Name: "tom",
		Age:  4,
	}
	catchan <- catA
	close(catchan)
	outchan := <-catchan
	fmt.Printf("outchan=%T,outchan=%v\n", outchan, outchan)
	fmt.Printf("outchan.Name=%v\n", outchan.Name)

	// 遍历
	intchan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intchan <- i * 2
	}
	close(intchan)
	fmt.Println("intchan长度： ", len(intchan))

	for v := range intchan {
		fmt.Println("管道中的值：", v)
	}
}
