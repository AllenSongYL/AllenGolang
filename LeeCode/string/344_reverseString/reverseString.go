package main

import "fmt"

// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

func reverseString(s []byte) {
	for start, end := 0, len(s)-1; start < len(s)/2; start++ {
		s[start], s[end] = s[end], s[start]
		end--
	}
}

func main() {
	astring := "xfasdfagsdgh"
	abyte := []byte(astring)
	reverseString(abyte)
	fmt.Println(string(abyte))

}
