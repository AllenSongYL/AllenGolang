# ants库
ants是一个高性能的 goroutine 池，实现了对大规模 goroutine 的调度管理、goroutine 复用，允许使用者在开发并发程序的时候限制 goroutine 数量，复用资源，达到更高效执行任务的效果。

https://github.com/panjf2000/ants/blob/master/README_ZH.md

检查当前 Worker 队列中是否有可用的 Worker，如果有，取出执行当前的 task；
没有可用的 Worker，判断当前在运行的 Worker 是否已超过该 Pool 的容量：{是 —> 再判断工作池是否为非阻塞模式：[是 ——> 直接返回 nil，否 ——> 阻塞等待直至有 Worker 被放回 Pool]，否 —> 新开一个 Worker（goroutine）处理}；
每个 Worker 执行完任务之后，放回 Pool 的队列中等待。

🚀 功能：
自动调度海量的 goroutines，复用 goroutines
定期清理过期的 goroutines，进一步节省资源
提供了大量有用的接口：任务提交、获取运行中的 goroutine 数量、动态调整 Pool 大小、释放 Pool、重启 Pool
优雅处理 panic，防止程序崩溃
资源复用，极大节省内存使用量；在大规模批量并发任务场景下比原生 goroutine 并发具有更高的性能
非阻塞机制


## 自定义池
ants支持实例化使用者自己的一个 Pool ，指定具体的池容量；通过调用 NewPool 方法可以实例化一个新的带有指定容量的 Pool ，如下：

// 设置10000个goroutine容量的池
p, _ := ants.NewPool(10000)

## 任务提交
提交任务通过调用 ants.Submit(func())方法：

ants.Submit(func(){})

## 动态调整 goroutine 池容量
需要动态调整 goroutine 池容量可以通过调用Tune(int)：

pool.Tune(1000) // Tune its capacity to 1000
pool.Tune(100000) // Tune its capacity to 100000

## 预先分配 goroutine 队列内存
ants允许你预先把整个池的容量分配内存， 这个功能可以在某些特定的场景下提高 goroutine 池的性能。比如， 有一个场景需要一个超大容量的池，而且每个 goroutine 里面的任务都是耗时任务，这种情况下，预先分配 goroutine 队列内存将会减少不必要的内存重新分配。

// ants will pre-malloc the whole capacity of pool when you invoke this function
p, _ := ants.NewPool(100000, ants.WithPreAlloc(true))

## 释放 Pool
pool.Release()

## 重启 Pool
// 只要调用 Reboot() 方法，就可以重新激活一个之前已经被销毁掉的池，并且投入使用。
pool.Reboot()

## 关于任务执行顺序
ants 并不保证提交的任务被执行的顺序，执行的顺序也不是和提交的顺序保持一致，因为在 ants 是并发地处理所有提交的任务，提交的任务会被分派到正在并发运行的 workers 上去，因此那些任务将会被并发且无序地被执行。

