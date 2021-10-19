// MIT License

// Copyright (c) 2018 Andy Pan

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ants

import (
	"errors"
	"log"
	"math"
	"os"
	"runtime"
	"time"
)

// 声明常量
const (
	// DefaultAntsPoolSize 设置默认goroutine 池的默认容量;\n
	// DefaultAntsPoolSize is the default capacity for a default goroutine pool。
	DefaultAntsPoolSize = math.MaxInt32

	// DefaultCleanIntervalTime 清理 goroutine 的间隔时间
	// DefaultCleanIntervalTime is the interval time to clean up goroutines.
	DefaultCleanIntervalTime = time.Second
)

const (
	// OPENED 表示池已打开。
	// iota 常量计数器。它是从零开始的索引。
	// OPENED represents that the pool is opened.
	OPENED = iota

	// CLOSED = 1
	// CLOSED represents that the pool is closed.
	CLOSED
)

var (
	// Error types for the Ants API.
	//---------------------------------------------------------------------------

	// ErrInvalidPoolSize will be returned when setting a negative number as pool capacity, this error will be only used
	// by pool with func because pool without func can be infinite by setting up a negative capacity.
	// 将池容量设置为负数时,将返回此错误
	// 将字符串包装成一个 error 对象返回
	ErrInvalidPoolSize = errors.New("invalid size for pool")

	// ErrLackPoolFunc will be returned when invokers don't provide function for pool.
	// 当调用者不为池提供函数时将返回ErrLackPoolFunc。
	ErrLackPoolFunc = errors.New("must provide function for pool")

	// ErrInvalidPoolExpiry will be returned when setting a negative number as the periodic duration to purge goroutines.
	// goroutine生存周期设置为负数时，返回ErrInvalidPoolExpiry
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	// ErrPoolClosed will be returned when submitting task to a closed pool.
	// 将任务提交到关闭的池时将返回ErrPoolClosed。
	ErrPoolClosed = errors.New("this pool has been closed")

	// ErrPoolOverload will be returned when the pool is full and no workers available.
	// 当池已满且没有可用的工作人员时，将返回ErrPoolOverload。
	ErrPoolOverload = errors.New("too many goroutines blocked on submit or Nonblocking is set")

	// ErrInvalidPreAllocSize will be returned when trying to set up a negative capacity under PreAlloc mode.
	// PreAlloc模式(预先把整个池的容量分配内存)下设置负容量，将返回ErrInvalidPreAllocSize
	ErrInvalidPreAllocSize = errors.New("can not set up a negative capacity under PreAlloc mode")

	//---------------------------------------------------------------------------

	// workerChanCap determines whether the channel of a worker should be a buffered channel
	// to get the best performance. Inspired by fasthttp at
	// https://github.com/valyala/fasthttp/blob/master/workerpool.go#L139
	// 设置工作管道的容量，是否缓冲
	workerChanCap = func() int {
		// Use blocking channel if GOMAXPROCS=1.
		// This switches context from sender to receiver immediately,
		// which results in higher performance (under go1.5 at least).
		// runtime.GOMAXPROCS(int) 当int<0,不更改配置。返回默认逻辑核心数量
		// 单核心的情况下，使用无缓冲管道
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		// Use non-blocking workerChan if GOMAXPROCS>1,
		// since otherwise the sender might be dragged down if the receiver is CPU-bound.
		// 默认有缓冲
		return 1
	}()

	defaultLogger = Logger(log.New(os.Stderr, "", log.LstdFlags))

	// Init a instance pool when importing ants.
	// 导入ants时初始化一个实例池。
	defaultAntsPool, _ = NewPool(DefaultAntsPoolSize)
)

// Logger is used for logging formatted messages.
// Logger 用于记录格式化的消息。
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})
}

// Submit submits a task to pool.
// 将任务提交到池中。
func Submit(task func()) error {
	return defaultAntsPool.Submit(task)
}

// Running returns the number of the currently running goroutines.
// 运行返回当前运行的 goroutines 的数量。
func Running() int {
	return defaultAntsPool.Running()
}

// Cap returns the capacity of this default pool.
// 返回这个默认池的容量。
func Cap() int {
	return defaultAntsPool.Cap()
}

// Free returns the available goroutines to work.
// 返回可用的 goroutines 数量。
func Free() int {
	return defaultAntsPool.Free()
}

// Release Closes the default pool.
// 关闭默认池。
func Release() {
	defaultAntsPool.Release()
}

// Reboot reboots the default pool.
// 重启默认池。
func Reboot() {
	defaultAntsPool.Reboot()
}
