# Go 流程控制

## **if判断**

注意

Go 语言中的 if 语句后面的条件判断表达式，不需要也不能加小括号，即 `()`，这是 Go 语言与其他语言 if 语句的区别。

同时，Go 语言 if 语句后面的大括号必须跟条件表达式写在一行，不能换行写，换行写会编译错误。

```
a := 6
b :=10
if b > a{
   fmt.Println("Yes!")
}
```

```
a := 10
b := 10
if a > b{
   fmt.Println("a big!")
} else if a < b {
   fmt.Println("b big!")
} else {
   fmt.Println("same")
}
```

### if 语句特殊写法

if 还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，代码如下：

~~~
if err := Connect(); err != nil {
    fmt.Println(err)
    return
}
~~~

Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。

err != nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。

这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if、else 语句组合中。



## for循环

go语言只有for循环，没有while循环。

格式

```
for ti := 10; ti < 100; ti++ {
   fmt.Println(ti)
}
```



死循环

```
a := 0
for {
   fmt.Println("haiCoder: ", a)
   a ++
   time.Sleep(time.Duration(3)*time.Second)
}
```

该 for 循环没有任何的开始条件和结束条件，因此是一个死循环。我们在循环开始之前定义了一个变量 i，并赋值为 0，每次执行 for 循环之后，i 的值加 1，同时每执行完一次，使用 Sleep() 函数，停止三秒钟。

## for range

**循环取出key：value**

使用for range的语法格式。可以用来遍历 **[字符串](https://haicoder.net/golang/golang-string.html)**、**[数组](https://haicoder.net/golang/golang-array.html)**、**[切片](https://haicoder.net/golang/golang-slice.html)**、**[map](https://haicoder.net/golang/golang-map.html)** 以及 **[channel](https://haicoder.net/golang/golang-channel.html)** 等。

### 循环取字符串

for range循环字符串

```
a := "ascdefg"

for keys, values := range a{
   fmt.Printf("keys index: %d --->key values: %v\n",keys ,values)
}

// 结果
keys index: 0 --->key values: 97
keys index: 1 --->key values: 115
keys index: 2 --->key values: 99
keys index: 3 --->key values: 100
keys index: 4 --->key values: 101
keys index: 5 --->key values: 102
keys index: 6 --->key values: 103
```



可以将keys换成 "_" 下划线，从而忽略key

```
for _, values := range a{
   fmt.Printf("keys index:  --->key values: %v\n",values)
}
```



只接收一个的时候，只返回字符串的索引，而不是字符串的值

```
for values := range a{
   fmt.Printf("print  %v\n",values)
}
// 结果
print  0
print  1
print  2
print  3
print  4
print  5
print  6
```



## switch分支结构

基于不同条件执行不同的行动，每一个case分支都是唯一的，从上到下逐一测试，直到匹配为止

匹配项的后面不需要加break

格式：

switch var1 {

​	case val1, val2:

​		...

​	case val3, val4:

​		...

​	default:

​		...

}



```
score := 80
	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
```



如果执行完匹配的case后，还需要继续执行后面的case，可以使用fallthrough

支持多个 case 连写的形式，即如果多个 case 执行相同的情况，则可以使用逗号分隔，写在一行。

switch var1 {    

​    case val1, val2:       

​        ...    

​    case val3, val4:        

​        ...    

​    default:        

​    ... 

}



## Break 语句

终止循环继续运行使用的关键字break

```
for i := 99; i < 200; i++{
   fmt.Println(i)
   if i > 121 {
      fmt.Println(i)
      break
   }
}
```

#### break终止指定循环

```
loop:
for i := 99; i < 133; i++{
   fmt.Println(i)
   if i > 111 {
      fmt.Println(i)
      break loop
   }
}
```



## Continue 语句

跳过本次循环，继续执行下一次循环。

Go 语言的 continue 语句，有与其他编程语言类似的功能，但不同的是，Go 语言的 continue 语句还可以选择具体跳过的循环。

注意 continue 与 **[break](https://haicoder.net/golang/golang-break.html)** 的区别，break 是直接终止了当前的循环，当前的循环不会再运行，而 continue 只是跳过本次循环，当前循环的后续循环还会继续运行。

~~~
loop:
for i := startIndex; i < endIndex; i++{
    //do something
    if condition {
        continue loop
    }
}
~~~

上面的循环是在 condition 条件满足的情况下，跳过 for 循环，运行 loop 同级层次的代码，注意这里的 loop 是标签，后面需要加上 `:`。

## goto 语句

在 **[Go 语言](https://haicoder.net/golang/golang-tutorial.html)** 中，实现 **[循环](https://haicoder.net/golang/golang-for.html)** 的跳转除了使用 **[break](https://haicoder.net/golang/golang-break.html)** 和 **[continue](https://haicoder.net/golang/golang-continue.html)** 以外，还可以使用 **[goto 语句](https://haicoder.net/golang/golang-goto.html)**。当然，goto语句不仅可以使用在循环中，还可以使用在代码的任何地方。

goto 语句使用最多的场景就是程序的错误处理，也就是当程序出错时，统一跳转到相应的标签处，统一处理错误。

语法

~~~
goto label

label:
    //do something
~~~

使用 goto 后面加上 label 名，可以直接将代码跳转到 label 的地方执行。

~~~
package main
import "fmt"
func main() {
	fmt.Println("start")
	//使用goto语句，跳出循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if j == 2{
				goto over
			}
		}
	}
	over:
		fmt.Println("Over")
}
~~~



## return 语句

在 **[Go 语言](https://haicoder.net/golang/golang-tutorial.html)** 中，跳转控制语句除了有 **[goto](https://haicoder.net/golang/golang-goto.html)** 、**[break](https://haicoder.net/golang/golang-break.html)** 和 **[continue](https://haicoder.net/golang/golang-continue.html)** 之外，还可以使用 **[return](https://haicoder.net/golang/golang-func-return.html)** 。

如果 return 语句使用在普通的 **[函数](https://haicoder.net/golang/golang-func.html)** 中，则表示跳出该函数，不再执行函数中 return 后面的代码，可以理解成终止函数。

如果 return 语句使用在 main 函数中，表示终止 main 函数，也就是终止程序的运行。

格式：

~~~
func fun(){
    //do something
    return
}

// 使用 return 语句，终止函数 fun 的执行。
~~~

~~~
package main
import "fmt"
func checkUserParam(val int){
	if val <= 0{
		fmt.Println("UserParam check Error")
		return
	}
	fmt.Println("UserParam check OK")
}
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	//使用return语句，终止函数执行
	checkUserParam(0)
	fmt.Println("Over")
}
~~~

