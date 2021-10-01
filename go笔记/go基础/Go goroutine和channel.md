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

P：  协程执行需要的上下文

G:    协程



### 运行状态

![image-20210930002332510](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20210930002332510.png)

![image-20210930002753077](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20210930002753077.png)



### 设置Golang运行的CPU数

go1.8 之后的版本，默认让程序运行在多核，可以不用设置

go1.8之前，需要配置，可以更高效的利用CPU

~~~
// runtime 包的NumCPU() 获取cpu数
	num := runtime.NumCPU()
	fmt.Println("CPU数：", num)
	// 设置最大数量
	runtime.GOMAXPROCS(num)
~~~





# 

## 代码案例

~~~
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
	results  = make(map[uint64]uint64, 10)
	// 声明一个全局的互斥锁
	// lock是一个全局的互斥锁
	// sync 是包
	// Mutex 互斥
	lock sync.Mutex
)

func Jx(i uint64) {
	var res uint64 = 1
	var a uint64
	for a = 1; a <= i; a ++ {
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


~~~



![image-20210930143231398](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20210930143231398.png)

多个协程，同时访问地址并写入造成问题

### 不同goroutine之间如何通讯

- 全局变量加锁同步
- channel

### 使用全局变量加锁同步

- 因为没有对全局变量加锁，所以会出现资源争夺问题，代码会出现错误
- 解决方案，加入互斥锁

# 管道channel

## 基本介绍

### 为什么使用channel

前面使用全局变量加锁同步来解决goroutine的通讯，但并不完美

- 主线程在等待所有goroutine全部完成的时间很难确定

- 如果主线程休眠时间长了，会增长等待时间，如果等待时间短了，可能还有goroutine处于工作状态，这是也会随主线程的退出而销毁

- 通过全局变量加锁同步来实现通讯，也并不利用多线程对全局变量的读写操作

  

### channel的介绍

- channel本质就是一个数据结构-队列
- 数据是先进先出 【FIFO first in first out】  
- 线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
- channel是由类型的，一个string的channel只能存放string类型数据



### 定义/声明channel

var 管道名称 管道类型 数据类型

var inchan  chan int    

var mapchan chan map[int]string 

var perchan chan Person
var perchan chan *Person



-  channel是引用类型
- channel必须初始化才能写入数据，即make后才能使用
- 管道是有类型的，intchan只能写入整数int

~~~
// 使用管道
	var intChan chan int
	intChan = make(chan int, 3)
	
	fmt.Println("intChan: ", intChan)
	fmt.Println(reflect.TypeOf(intChan))
	fmt.Printf("intChan的值: %v, intChan本身的地址：%p", intChan, &intChan)
// 输出
intChan:  0xc00001c100
chan int
intChan的值: 0xc0000d0080, intChan本身的地址：0xc0000ca018
~~~



### 写入管道

管道名 <- 值

```
// 向管道写入数据
intChan <- 10
num := 211
intChan <- num
```

注意：

- 向管道中写入数据通常会导致程序阻塞，直到由其他goroutine从这个端到中读取数据

- 当我们向管道写入数据时，不能超过其容量



### 读取数据

var num2 int

num2，ok = <-intchan

取完后，intchan中的值被取出就不存在了

全部取出再取，会报错 deadlock

~~~

type cat struct {
	Name string
	Age int
}

// 创建一个cat结构体管道
	catchan := make(chan cat, 10)
	catA := cat{
		Name: "tom",
		Age: 4,
	}
	catchan <- catA
	outchan := <- catchan
	fmt.Printf("outchan=%T,outchan=%v", outchan, outchan)
	// 输出：outchan=main.cat,outchan={tom 4}
	fmt.Printf("outchan.Name=%v\n", outchan.Name)
	// 输出：outchan.Name=tom
	
~~~

注意：如果管道中没有数据，会导致程序阻塞，直到有数据



### 管道的长度和容量

len(chan)     //长度

cap(chan)    //容量 

```
fmt.Printf("channel len=%v cap=%v \n", len(intChan),cap(intChan))
```



### 关闭管道

close(管道名)

使用内置函数close可以关闭管道，关闭后，无法再向channel写入数据，但是任然可以读取数据



### 管道的遍历

channel支持for-range的方式进行遍历

- 遍历时，如果channel没有关闭，则会出现deadlocak的错误
- 遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成后，退出遍历
- 遍历管道不能使用普通的for循环，因为管道的长度会变化

~~~
// 遍历
	intchan := make(chan int,100)
	for i := 0; i < 100; i ++ {
		intchan <- i *2
	}
	close(intchan)
	fmt.Println("intchan长度： ",len(intchan))

	for v := range intchan {
		fmt.Println("管道中的值：", v)
	}
~~~



### 管道阻塞

如果只向管道写入数据，而没有读取，就会出现阻塞而dead lock，原因是往里写的数据超过了管道容量。

编译器运行，发现一个管道只有写没有读，则该管道会阻塞。

写管道和读管道的频率不一致，无所谓。



### 只读只写

管道可以声明为只读或者只写

- 在默认情况下，管道时双向的，可读可写

  ~~~
  var chan1 chan int
  ~~~

- 只写

  ~~~
  var chan2 chan <- int
  ~~~

- 只读

  ~~~
  var chan3  <- chan int
  ~~~

  

# select

传统的方法遍历管道时，如果不关闭会阻塞而导致dead lock

可以使用select解决从管道取数据的阻塞问题

~~~
for {
	select {
// 这里，如果intchan一直没有关闭，不会一直阻塞而报错
// 会自动到下一个case匹配
	case v := intchan :
		fmt.Println("", v)
	case v := <-stringchan :
		fmt.Println("", v)
	default:
		fmt.Println("可以加入自己的逻辑")
		// 结束
		return
	}
}

~~~



goroutine中使用reco

