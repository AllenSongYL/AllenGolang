# GO

## 第一个程序

~~~
t2.go

// package main定义了包名。必须在源文件中非注释的第一行指明这个包属于哪个包
// package表示一个可独立执行的程序，每个GO应用程序都包含一个名为main的包
package main

// 告诉编译器这个程序需要使用fmt包 fmt包实现了格式化IO（输入/输出）的函数
import "fmt"

/* 多行注释 */
// 单行注释
func main() {s
  fmt.Println("Hello, World!")
}
~~~



go build  t2.go

// 生成二进制文件。 生成了一个t2.exe



fmt.Println()

// 将字符串输出到控制台，并在最后自动增加换行字符 \n

fmt.Print("hello, world\n")     可以得到相同的结果

**注意**

~~~
fun main()
{

	fmt.Prinln("helloworld")

}
~~~

