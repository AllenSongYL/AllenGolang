# Go 语言容器



**[Go 语言](https://haicoder.net/golang/golang-tutorial.html)** 为开发者提供了内置的四种常用数据结构：**[数组](https://haicoder.net/golang/golang-array.html)**、**[切片（slice）](https://haicoder.net/golang/golang-slice.html)**、**[列表(list)](https://haicoder.net/golang/golang-list.html)** 以及 **[字典（map）](https://haicoder.net/golang/golang-map.html)** 用来保存一组数据。

数组、切片（slice）以及字典（map）这三种数据结构都是用于同时保存多个数据项。

数组和切片（slice）比较相似，它们都按顺序保存元素相同类型的元素，每个元素都有自己的索引，因此数组和切片都可通过索引访问元素。

字典（map）存储的数据都是无序的，其中字典是用 `key-value` 的形式保存数据。

### 



## 数组

go语言中的数组是由固定长度的特定类型组成，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，所以在golang中很少使用数组，一般都是使用 **[切片](https://haicoder.net/golang/golang-slice.html)** 来代替数组。

var  varname  [count]type



索引位从0开始。可以通过索引修改值，所以数组时**可变的**

定义一个数组 varName，该数组拥有 count 个元素，每个元素的类型都是 Type。

```
var ashu [3]int
ashu[0] = 2
ashu[1] = 3
ashu[2] = 4
fmt.Println(ashu)
//结果
[2 3 4]
```

**获取数组的类型**

reflect.TypeOf()

```
fmt.Println("ashu type:", reflect.TypeOf(ashu)
```



## Go语言数组初始化二

数组初始化一

~~~
var varName [count]Type = [count]Type{element1, element2, element3}

我们定义了一个数组 varName，该数组有 count 个元素，每个元素的类型都是 Type，同时，我们使用 element1、element2、element3 来初始化了该数组。
~~~



数组初始化二

~~~
var varName = [count]Type{element1, element2, element3}

我们定义了一个数组 varName，该数组有 count 个元素，每个元素的类型都是 Type，同时，我们使用 element1、element2、element3 来初始化了该数组。
~~~



数组初始化三

~~~
var varName = [...]Type{element1, element2, element3}

这里的数组的个数我们使用 …，即表示数组元素的个数是根据元素个数自动推导。
~~~



数组初始化四

~~~
var varName = [...]Type{index0:element1, index1:element2, index2:element3}

这里我们给索引 index0 的位置设置值为 element1，index1 的位置设置值为 element2，index2 的位置设置值为 element3。
~~~



## 数组赋值

给数组arr的的索引位index的位置设置值为value

~~~
package main
import (
	"fmt"
)
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	//给定义好的数组的指定索引位置处赋值
	var arrHaiCoder [3]string
	arrHaiCoder[0] = "Hello"
	arrHaiCoder[1] = "嗨客网"
	arrHaiCoder[2] = "HaiCoder"
	fmt.Println("arrHaiCoder0 =", arrHaiCoder[0])
	fmt.Println("arrHaiCoder1 =", arrHaiCoder[1])
	fmt.Println("arrHaiCoder2 =", arrHaiCoder[2])
}
~~~

### 重新赋值

数组指定索引的位置有值后，也可以通过索引重新设置值

### 注意事项

1. 数组时多个相同类型数据的组合，一个数组一旦声明，其长度时固定的，不能动态的变化
2. var arr  []int是一个slice切片   中括号中不写
3. 数组中的元素可以是任何数据类型，包括值类型和引用数据类型，但是不能混用
4. 数组创建后，如果没有赋值，有默认值(零值，“”， false)
5. 下标从0开始，必须在指定范围内使用，否则报panic数组越界
6. 数组为值类型，默认情况下的值传递，为值拷贝，数组间不会影响
7. 如果在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）





## 数组的遍历

1. for

2. for range



## 数组的比较

数组的比较使用 == 的方式，如果数组的元素个数不相同，那么不能比较数组

比较数组arr与数组arr1是否相等，如果相等，则返回true，反则，返回false

数组长度不相同，不可以通过 == 来比较数组





## 二维数组

~~~
var varName [count][count2]Type

~~~

| 数        | 描述                       |
| --------- | -------------------------- |
| *var*     | 定义数组使用的关键字。     |
| *varName* | 数组名。                   |
| *count*   | 二维数组的行数。           |
| *count2*  | 二维数组的列数。           |
| *type*    | 二维数组中每个元素的类型。 |







## 三维数组

~~~
var varName [count][count2][count3]Type
~~~

定义一个三维数组 varName，该数组的每一个元素都是一个二维数组，二维数组拥有 count2 行 count3 列，每个元素的类型都是 Type。







## 切片

golang中的切片是数组的引用，因此切片是引用类型

切片的使用和数组类似，遍历切片，访问切片的元素和求切片的长度len与数组都一样。但切片的**长度是可以变化的**，不像数组是固定的，因此也可以说切片是可以动态变化的数组。

可以使用**len()** 函数获取长度



varname  :=  []type{element1, element2}

定义一个切片varname，改切片每个元素的类型都是type，目前又两个元素

创建切片时，只需要指定切片的类型不需要指定切片的长度

可以使用reflect.TypeOf() 获取切片类型

###  创建切片

- 从数组创建

  var  sliceName = arr[index1: index2]

  | 参数        | 描述                       |
  | ----------- | -------------------------- |
  | *var*       | 定义切片变量使用的关键字。 |
  | *sliceName* | 切片变量名。               |
  | *arr*       | 数组名。                   |
  | *index1*    | 数组的开始索引。           |
  | *index2*    | 数组的结束索引。           |

  创建一个sliceName。该切片元素的内容是从数字arr的索引index1开始到索引index2结束。

  ~~~
  package main
  import (
  	"fmt"
  )
  func main() {
  	
  	//从已存在的数组的内容创建切片
  	var arrHaiCoder = [3]string{"Hello", "嗨客网", "HaiCoder"}
  	
  	var sliceHaiCoder = arrHaiCoder[1:3]
  	fmt.Println("sliceHaiCoder =", sliceHaiCoder)
  }
  ~~~

  

- **使用make创建**

  
  特点：
  
  1. 通过make()方式创建切片可以指定切片的大小和容量
  2. 如果没有给切片的各个元素赋值，那么就会使用默认值int,float=>0; string=>""; bool=> false
  3. 通过make()方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素
  
  
  
  var sliceName []type = make([]type,len,[cap])

| 参数        | 描述                       |
| ----------- | -------------------------- |
| *var*       | 定义切片变量使用的关键字。 |
| *sliceName* | 切片变量名。               |
| *type*      | 切片的每一个元素的类型。   |
| *len*       | 切片的长度。               |
| *cap*       | 可选，切片的容量。         |







~~~
package main
import (
	"fmt"
)
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	//使用make指定切片的长度和容量创建切片
	var sliceHaiCoder = make([]string, 3, 3)
	//创见一个长度为3容量为3的切片，接着使用索引的方式给切片赋值
	
	sliceHaiCoder[0] = "Hello"
	sliceHaiCoder[1] = "嗨客网"
	sliceHaiCoder[2] = "HaiCoder"
	fmt.Println("sliceHaiCoder =", sliceHaiCoder)
}
~~~





- **指定数组创建**

定义一个切片，直接指定具体数组，使用原理和make类似

var sliceName  []string = []string{"","",""}



### 切片遍历

- for

- for  range

和数组一致





## 字典map

map时key-value数据结构。可以通过key来快速检索。





### 创建map

~~~

// 第一张 
// 先声明,需要先make，make的作用是给map分配数据空间
var a map[string]string
a = make(map[string]string, 10)

// 第二章
// 声明后使用make 
var cities = make(map[string]string, 10)
var cities2 = make(map[string]string)

// 声明直接赋值
var Name map[string]string = map[string]string{
    "no1" : "wuxi"
}
// 使用推导式，最后别要忘了逗号
varName  :=  map[Type]Type2{
	"hero" : "sx",
	"xxx" : "rrr",
}
~~~



不初始化map，会创建一个nil map。nil map不能用来存放键值对。

定义一个字典 varName，该字典的 Key 的类型为 Type1，Value 的类型为 Type2。

### 增加更新

map["key"] = value

如果key没有则添加，如果key存在则修改



### 删除元素

**delet()函数**

delete(mapName, "key")

delete()函数用于删除集合的元素，参数为map和其对应的key

如果为nil则不进行任何操作

要删除所有的key，遍历key

或者make一个新的，让原来的被gc回收



### 查找元素

findRes存在key则返回true，否则返回false



### 遍历map

for k,v := range cities {

​    fmt.Printf("k=%v,v=%v", k, v)

}

### map长度

len(Mapname)



### map切片

~~~
var a []map[string]string
a = make([]map[sring]string, 2)

if a[0] == nil {
	a[0] = make(map[string]string, 2)
	a[0]["name"] = "xx"
	a[0]["age"] = "16"
// 但是这种方法超出范围会报panic 
// 这里我们使用切片的append函数,可以动态的增加
}

newa := map[string]string{
   xxx : xxx,
}
a = append(a,newa)
~~~









### sync.map

go中map如果在并发读的情况下是线程安全的，如果是在并发写的情况下，线程是不安全的。Golang 为我们提供了一个 sync.Map 是并发写安全的。





## 结构体

go语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型

结构体是由一系列具有相同类型和不同类型的数据构成的数据集合

### 定义结构体

结构体定义需要使用type和struct语句。struct语句定义一个新的数据类型，结构体中有一个或多个成员。type语句设定了结构体的名称。结构体的如下

~~~
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
~~~

一旦定义了结构体，就能用于变量的声明，格式如下

~~~
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
~~~

### 访问结构体成员

如果要访问结构体成员，需要使用点“.”操作符，格式为

~~~
结构体.成员体
~~~

~~~
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}
~~~



### 结构体作为函数参数

可以想其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量。

~~~
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
}

func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}
~~~



### 结构体指针

格式

~~~
var struct_pointer *books
~~~

以上定义的指针变量可以存储结构体**变量的地址**。查看结构体变量，可以将&符号放置于结构体变量前

~~~
struct_pointer = &book
~~~

使用结构体指针访问结构体成员，使用“.”操作符

struct_pointer.title

结构体中属性的首字母大小写

首字母大写相当于public（公有的）

首字母小写相当于private（私有的）

**注意点**

当要将结构体对象转换为JSON时，对象中的属性首字母必须是大写，才能正常转换为JSON。

# 语言类型转化

类型转换用于将一种数据类型的变量转换为另一种类型的变量

格式

type_name(e)

例子

var sum int =17

float32(sum) 

