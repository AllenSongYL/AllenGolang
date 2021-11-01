package main

import "fmt"

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

// 我的代码
func replaceSpace(s string) string {
	var results string
	for _, value := range s {
		svalue := string(value)
		if svalue == " " {
			svalue = "%20"
		}
		results = results + svalue
	}
	return results
}

// 原地修改.参考答案
func replaceSpace2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}

func main() {
	//fmt.Println(replaceSpace("x yy zz cc"))
	fmt.Println(replaceSpace2("x yy zz cc"))
}
