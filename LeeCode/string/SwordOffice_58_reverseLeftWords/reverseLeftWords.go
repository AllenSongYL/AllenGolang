package main

import "fmt"

//func reverseLeftWords(s string, n int) string {
//	results:=make([]string, len(s))
//	for i:=0;i<len(s);i++{
//		results[i] = string(s[i])
//		//tmp := string(s[i])
//		//s = s[i+1:] + tmp
//	}
//	fmt.Println(results)
//	for  i:=0;i<n;i++{
//
//	}
//	return s
//}

func reverseLeftWords(s string, n int) string {
	s = s[n:] + s[:n]
	return s
}

func main() {
	test := "abcdefg"
	fmt.Println(reverseLeftWords(test, 3))
}
