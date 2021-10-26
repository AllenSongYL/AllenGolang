package main

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
// 输入：x = 123   x = -123    120
// 输出：321       -321        21

import (
	"fmt"
	"time"
)

func reverse(x int) int {
	var res int
	for ; x != 0; x /= 10 {
		// 如果溢出则返回0
		if temp := int32(res); (temp*10)/10 != temp {
			fmt.Println("zhi", (temp*10)/10)
			return 0
		}
		res = res*10 + x%10
		// 通过除以10，逐步拿到x的所有位，5 / 10 = 0 退出循环
	}
	return res
}

func main() {
	start := time.Now()
	testInt := 123456789
	fmt.Println(reverse(testInt))
	fmt.Println(time.Since(start))
}
