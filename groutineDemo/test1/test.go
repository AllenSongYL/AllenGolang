package main

// 带缓冲区的channel

import (
	"fmt"
)

func produce(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}

}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("Receive:", v)
	}
}

func main() {

	ch := make(chan int, 10)
	ticket := make(chan bool, 1)
	defer close(ticket)

	//var wg sync.WaitGroup
	//wg.Add(1)
	go func() {
		produce(ch)
		//defer wg.Done()
	}()

	//wg.Add(1)
	go func() {
		consumer(ch)
		defer func() {
			ticket <- true
		}()
		//defer wg.Done()
	}()

	//wg.Wait()
	<-ticket

}
