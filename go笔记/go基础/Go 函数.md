# Go     函数

为了完成某一功能的程序指令（语句）的集合，称为函数。**[Go 语言](https://haicoder.net/golang/golang-tutorial.html)** 的函数可以分为：自定义函数和系统函数。

Go 语言函数与其他语言函数最大的不同是，Go 语言的函数可以支持 **[返回任意多个值](https://haicoder.net/golang/golang-func-return.html)**，而其他语言的函数一般只支持返回一个值。

Go 语言的函数也支持普通函数、**[匿名函数](https://haicoder.net/golang/golang-anonymous-func.html)** 和 **[闭包](https://haicoder.net/golang/golang-closure.html)** 三种形式。

在程序中，编写函数的主要目的是将一个需要很多行代码的复杂问题分解为一系列简单的任务来解决，而且，同一个任务（函数）可以被多次调用，有助于代码重用。

```
 • 无需声明原型。
    • 支持不定 变参。
    • 支持多返回值。
    • 支持命名返回参数。 
    • 支持匿名函数和闭包。
    • 函数也是一种类型，一个函数可以赋值给变量。

    • 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
    • 不支持 重载 (overload) 
    • 不支持 默认参数 (default parameter)。
```



语法

~~~
func funcName(paramlist paramType)(returnval returnType){
	// 执行语句...
	return valuelist
}
~~~



| 参数         | 描述                     |
| ------------ | ------------------------ |
| *func*       | 定义函数所使用的关键字。 |
| *funcName*   | 函数名。                 |
| *paramlist*  | 函数参数列表。           |
| *paramType*  | 函数参数类型。           |
| *returnval*  | 可选，函数返回值。       |
| *returnType* | 函数返回值类型。         |
| *return*     | 函数返回值使用的关键字。 |
| *valuelist*  | 函数返回值列表。         |

函数可以有返回值，也可以没有返回值。同时，Go 语言函数也支持返回多个值。

### 例子

~~~
package main
import "fmt"
func sum(a, b int)int{
	return a+b
}
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	//用函数，实现计算任意两个数的和
	result := sum(10,20)
	fmt.Println("Result =", result)
}
~~~

### 值传递类型

- 值传递    指在调用函数时**将实际参数复制一份传递到函数中**，这样在函数中如果对参数进行修改，将不会影响到实际参数。

~~~
func add(nums2,nums3 int)(nums4 int ){
	nums2 = nums2 + nums2
	nums4 = nums2 + nums3
	return nums4
}

fun main(){
	var num1 int = 10
	var num2 int = 22
	num3 := add(num1, num2)
	fmt.Println(num1)
	fmt.Println(num3)
}
~~~



- 引用传递    是指在调用函数时将**实际参数的地址传**递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

~~~
// *int 指针类型
func swap(a,b *int){
	var tmp int
	// 传入的a是指针地址，*a通过地址取值
	tmp = *a
	*a = *b
	*b =tmp
}


func main(){
	var a int = 3
	var b int =4
	swap(&a, &b)
	//& 获取指针地址

	fmt.Println(a)
	fmt.Println(b)

}
~~~

默认情况下，go语言使用的是值传递即在调用过程中不会影响到实际参数。

一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低

**map、slice、chan、指针、interface默认以引用的方式传递。**



### 不定参数

函数的参数不固定，后面的类型是固定的

可变参数本质上就是slice。只能有一个，且必须是最后一个。

有参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。

~~~
func myfunc(args ...int) {    //0个或多个参数
  }

 func add(a int, args…int) int {    //1个或多个参数
  }

 func add(a int, b int, args…int) int {    //2个或多个参数
  }
~~~



### 任意类型的不定参数

就是函数的参数和每个参数的类型都不是固定的。

用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。

~~~
func myfunc(args ...interface{}) {
}
~~~









## 返回值

1. 无返回值时，不需要使用return函数
2. 返回多个值

~~~
func funcName(param1, param2 paramType1, ...)(returnType1, returnType2, ...){
	//执行语句...
	return 返回值列表
}


~~~

3. “_” 标识符，用来忽略函数的某个返回值
4. Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
5. 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。

~~~
func add(a, b int) (c int) {
    c = a + b
    return
}
~~~



命名返回参数允许 defer 延迟调用通过闭包读取和修改。

~~~
package main

func add(x, y int) (z int) {
    defer func() {
        z += 100
    }()

    z = x + y
    return
}

func main() {
    println(add(1, 2)) 
}
~~~

显式 return 返回前，会先修改命名返回参数。

~~~
package main

func add(x, y int) (z int) {
    defer func() {
        println(z) // 输出: 203
    }()

    z = x + y
    return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
    println(add(1, 2)) // 输出: 203
}
~~~



## init函数

每一个源文件都可以包含一个init函数，该函数会在main函数执行之前，被go框架调用。

通常可以在init函数中完成初始化工作

1. 如果一个文件同时包含全局变量定义，init函数和main函数，先执行的流程是 先全局变量的定义 ---> init函数 ---> main函数
2. 





## 匿名函数

匿名函数是指不需要定义函数名的一种函数实现方式。

匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

~~~
package main

import (
    "fmt"
    "math"
)

func main() {
    getSqrt := func(a float64) float64 {
        return math.Sqrt(a)
    }
    fmt.Println(getSqrt(4))
}
~~~

用法

1. 在定义匿名函数时直接调用

~~~
res1 := func(n1 int, n2 int) int {
	return n1 + n2
}(10, n2)
~~~



2. 将匿名函数赋值给一个变量（函数变量），再通过改变量来调用匿名函数





## 闭包 递归

闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。

“官方”的解释是：所谓“闭包”，指的是一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也是该表达式的一部分。

例子

~~~
package main

import (
    "fmt"
)

func a() func() int {
    i := 0
    b := func() int {
        i++
        fmt.Println(i)
        return i
    }
    return b
}

func main() {
    c := a()
    c()
    c()
    c()

    a() //不会输出i
}
~~~

~~~
package main

import (
	"fmt"
)

func main() {
	a:=Addupper()
	fmt.Println(a(1))
	fmt.Println(a(1))
	fmt.Println(a(1))
	fmt.Println(a(1))
}

// Addupper() 函数，返回的数据类型是fun (int) int
// 闭包的说明：返回的是一个匿名函数，但是这个函数引用到函数外的n
// 这个匿名函数和n形成一个整体，构成闭包
// 大家可以这么理解：闭包是一个类，而函数是操作，n是字段
// 当反复的调用a时，n只初始化一次，因此每调用一次就进行累加
func Addupper() func (int) int {
	var n int = 10
	var str string = "hello"
	return func (x int) int {
		n = n + x
		str += "x"
		fmt.Println(str)
		return n
	}
}
~~~





闭包复制的是原对象指针，这就很容易解释延迟引用现象。



## 递归函数

一个函数调用自己

~~~
package main

import "fmt"

func factorial(i int) int {
    if i <= 1 {
        return 1
    }
    return i * factorial(i-1)
}

func main() {
    var i int = 7
    fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}
~~~

斐波那契数列

~~~
package main

import "fmt"

func fibonaci(i int) int {
    if i == 0 {
        return 0
    }
    if i == 1 {
        return 1
    }
    return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
    var i int
    for i = 0; i < 10; i++ {
        fmt.Printf("%d\n", fibonaci(i))
    }
}
~~~



## 延迟调用



### defer特性

1. 关键字defer用于注册延迟调用
2. 这些调用知道return之前才被执行。因此可以用来做资源清理
3. 多个defer语句，按先进后出的方式执行
4. defer语句中的变量，再defer声明时就决定了

### defer用途

1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

## 异常处理

- **panic 抛出错误**

  1. 内置函数
  2. 加入函数中书写了panic语句，会终止其后面要执行的代码，再panic所在函数内如果存在defer函数列表，按照defer的逆序执行
  3. 返回函数的调用者a，在a中，调用函数语句之后的代码不会执行。加入函数a中存在要执行的defer函数列表，按照defer的逆序执行
  4. 知道goroutine整个退出，并报告错误

  

- **recover 捕获错误**

  1. 内置函数

  2. 用来控制一个goroutine的panicking行为，捕获panic，从而影响后续执行

  3. 一般的建议

     a：在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行

     b：可以获取通过panic传递的error

1. 利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
2. recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
3. 多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

