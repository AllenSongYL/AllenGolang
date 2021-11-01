package main

import "fmt"

// 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
//返回所有可能的结果。答案可以按 任意顺序 返回。

//示例 1：
//输入：s = "()())()"
//输出：["(())()","()()()"]

//示例 2：
//输入：s = "(a)())()"
//输出：["(a())()","(a)()()"]

//示例 3：
//输入：s = ")("
//输出：[""]

func removeInvalidParentheses(s string) []string {
	slen := len(s)
	fmt.Println("string 长度： ", len(s))
	res := make([]string, 0, slen)
	for i := 0; i < slen; i++ {
		res = append(res, string(s[i]))
	}
	return res
}

func main() {
	test := "\"{()}\"[]a(()"
	res := removeInvalidParentheses(test)
	fmt.Println(res)
	fmt.Println(len(res))
}
