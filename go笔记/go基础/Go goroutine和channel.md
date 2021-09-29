# goroutine

## 基本介绍

###  进程和线程说明

- 进程就是程序在操作系统中一次执行的过程，时系统进行资源分配和调度的基本单位

- 线程是进程的一个执行实例，是程序执行的最小单元，它是比进程更小的能独立运行的基本单位

- 一个进程可以创建和销毁多个线程，同一进程中的多个线程可以并发执行

- 一个程序至少有一个进程，一个进程至少有一个线程

  

### 并发和并行

  - 对线程在单核上运行，就是并发

    作用在一个CPU

    在一个时间点上，只有一个任务在执行

    

  - 多线程在多核上运行，就是并行

    多个任务作用在多个CPU

    一个时间点上，由多个任务在同时执行

    并行速度比并发快

### GO协程和GO主线程

- GO主线程（可以理解为进程）： 一个Go线程上，可以起多个协程，协程是轻量级的线程
- Go协程的特点
  - 有独立的栈空间
  - 共享程序栈空间
  - 调度由用户控制
  - 协程是轻量级的线程



### 例子

输出效果说明，main这个主线程和test协程同时执行

~~~
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


###
程序开始时间： 2021-09-30 00:06:52
mian(): Hello,Golang 1
HelloWorld():Hello,Golang 1
HelloWorld():Hello,Golang 2
mian(): Hello,Golang 2
mian(): Hello,Golang 3
HelloWorld():Hello,Golang 3
HelloWorld():Hello,Golang 4
mian(): Hello,Golang 4
mian(): Hello,Golang 5
HelloWorld():Hello,Golang 5
HelloWorld():Hello,Golang 6
mian(): Hello,Golang 6
mian(): Hello,Golang 7
HelloWorld():Hello,Golang 7
HelloWorld():Hello,Golang 8
mian(): Hello,Golang 8
HelloWorld():Hello,Golang 9
mian(): Hello,Golang 9
mian(): Hello,Golang 10
HelloWorld():Hello,Golang 10
结束时间： 2021-09-30 00:07:02
运行时长： 10.0308692s

进程 已完成，退出代码为 0
~~~



### 流程图

![image-20210930000935700](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20210930000935700.png)



### 注意点

- 如果主线程退出，即使协程还没执行完毕，也会退出。

- 协程可以在主程序没有退出前结束。
- 主线程是一个物理线程，直接作用在cpu上的。是重量级的非常耗费cpu资源。
- 协程从主线程开启，是轻量级的线程，是逻辑态，对资源消耗相对小
- golang的协程机制是重要的特点，可以轻松开启上万个协程。其他编程语言的开发机制一般是基于线程的，开启过多的线程，资源耗费大，这里就凸显了Golang在并发上的优势



## groutine的调度模型

### MPG模式基本介绍

M： 操作系统的主线程（是物理线程）

P： 协程执行需要的上下文

G:   协程
