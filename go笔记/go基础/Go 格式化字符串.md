# Go  类型转换



## 	1. 基本数据类型转字符串



### 		fmt.Sprintf方法

- 通用
  - **%v**              // 值的默认格式输出
  - %+v            // 类似%v,但输出结构体时会添加字段名
  - %#v            // 值的Go语法表示
  - **%T**             // 值的类型Go语法表示
  - %%            // 百分号
- 布尔值
  - **%t**             // 单词true或false
- 整数
  - %b             // 表示为二进制
  - %c             // 表示为unicode码值
  - %d             // 表示为十进制             
  - %o             // 表示为八进制
  - %q             // 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
  - %x             // 表示为十六进制，使用a-f
  - %X             // 表示为十六进制，使用A-F
  - %U            // 表示为Unicode格式：U+1234，等价于"U+%04X"
- 浮点与复数
  - %b            // 无小数部分、二进制指数的科学计数法，如-123456p-78
  - %e            // 科学计数法，如-1234.456e+78
  - %E            // 科学计数法，如-1234.456E+78
  - %f            // 有小数部分但无指数部分，如123.456
  - %F            // 等价于%f
  - %g            // 根据实际情况采用%e或%f格式
  - %G            // 根据实际情况采用%E或%F格式
- 字符串和byte
  - %s             // 直接输出字符串或者[]byte
  - %q            // 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
  - %x            // 每个字节用两字符十六进制数表示（使用a-f）
  - %X            // 每个字节用两字符十六进制数表示（使用A-F） 
- 指针
  - %p            // 表示为十六进制，并加上前导的0x   



```

import (
	"fmt"
)

func main() {
	num1 := 99
	str := ""
	str = fmt.Sprintf("num1: %d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)
	// str type string str=num1: 99
}
```



### 		strconv包



```
var num3 int = 77
var num4 float64 = 23.56
var b2 bool = true
str, str2, str3 := "", "", ""

str = strconv.FormatInt(int64(num3), 20)
fmt.Printf("str type %T str=%q\n", str, str)

str2 = strconv.FormatFloat(num4, 'f', 10, 64)
// fmt: 转成的格式'f'(ddd.dddd)  'b'(-dddd+-)
// prec： 小数点后保留位数
// bitSize: 表示数类型int64
fmt.Printf("str type %T str=%q\n", str2, str2)

str3 = strconv.FormatBool(b2)
fmt.Printf("str type %T str=%q\n", str3, str3)
```

strconv.Itoa(x)   

不能直接使用int64(),  strconv.Itoa(int64(x))   

```
var num5 int = 4567
var str string

str = strconv.Itoa(num5)
fmt.Printf("str type %T str=%q\n", str, str)
fmt.Printf(str)
```



## 	2. string转基本数据类型

### 		strconv.ParseBool(x)

​		这个函数会返回两个值，使用_忽略错误的返回值

​		字符串转bool

```
var str1 string = "true"
var b bool
b, _ = strconv.ParseBool(str1)
fmt.Printf("str type %T str=%t\n", b, b)
```





