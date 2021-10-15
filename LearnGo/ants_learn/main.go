package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	// atomic提供的原子操作能够确保同一时刻只有一个goroutine进行原子操作
	"sync/atomic"
	"time"
)

var sum int32

func Myfunc(i interface{}) {
	// 类型断言
	n := i.(int32)

	// 原子增。第一个参数必须是地址
	// Add前缀为增。
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	// 睡10毫秒
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello world demoFunc() sleep 10 millisecond")
}

func main() {

	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)

	// ants.Release() 关闭默认的池
	defer ants.Release()

	// 设置运行次数1000
	runTime := 1000

	var wg sync.WaitGroup

	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	// 将1000个任务提交到 任务池
	for i := 0; i < runTime; i++ {
		wg.Add(1)
		// 将任务提交到池中
		_ = ants.Submit(syncCalculateSum)
	}

	wg.Wait()
	// 运行返回当前运行的 goroutines 的数量。
	fmt.Printf("running groutine %d\n", ants.Running())
	fmt.Println("finish all tasks submit to ant pool~")

	// 使用一个函数池
	// 将 goroutine 池的容量设置为 10，将过期时间设置为 1 秒。
	// 生成具有特定功能的蚂蚁池实例。
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		Myfunc(i)
		wg.Done()
	})

	defer p.Release()
	for i := 10; i < runTime; i++ {
		wg.Add(1)
		// Invoke 将任务提交到池中
		_ = p.Invoke(int32(i))
	}

	wg.Wait()
	fmt.Println("running gorotines: %d\n", p.Running())
	fmt.Printf("finish all tasks,results is %d\n", sum)

	// 程序结束输出
	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)

}
