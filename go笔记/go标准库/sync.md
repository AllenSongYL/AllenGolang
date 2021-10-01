提供了基本同步，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些

## Mutex

互斥锁

可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。

 - Lock

   Lock 方法 锁住m，如果m已经加锁，则阻塞直到m解锁

   ~~~
   func (m *Mutex) Lock()
   ~~~

   

 - Unlock

   解锁