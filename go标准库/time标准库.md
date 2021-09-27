# Time标准库





time中定义的结构体

type Time struct {
    wall uint64
    ext  int64
    loc  *Location
}



## time.Now()

输入： 无
输出： time.Time(结构体)类型的
作用： 返回当前时间



~~~
timeNow := time.Now()
fmt.Println(timeNow)
fmt.Println(reflect.TypeOf(timeNow))
//2021-09-27 10:44:47.6943412 +0800 CST m=+0.004395901
//time.Time

~~~

## time.Time.Unix()

输入： Time
输出： int64的unix时间戳
作用： 转换为unix时间

# 实例方法



## time.Time转string

timeNow.String()

```
timeStringNow := timeNew.String()
fmt.Println(timeStringNow)
fmt.Println(reflect.TypeOf(timeStringNow))
//2021-09-27 10:46:49.3379938 +0800 CST m=+0.003894301
//string
```



## time.Time转Unix时间戳

timeNow.Unix()

```
t2 := timeNow.Unix()
fmt.Println(t2)
fmt.Println(reflect.TypeOf(t2))
//1632711410
//int64
```



## time.Time格式化字符串

timeNow.Foramt()

```
t2 := timeNew.Format("2006-01-02 15:04:05")
fmt.Println(t2)
fmt.Println(reflect.TypeOf(t2))
//2021-09-27 10:59:39
//string
```





```
timeStart := time.Now()
timeEnd := time.Now()
timeSub := timeEnd.Sub(timeStart)
fmt.Println("运行时长：", timeSub)
fmt.Println(reflect.TypeOf(timeSub))
// 运行时长： 227.5292ms
//time.Duration
```