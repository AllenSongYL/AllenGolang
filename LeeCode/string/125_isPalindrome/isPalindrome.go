package main

import (
	"fmt"
)

//func isPalindrome(s string) bool {
//	if s == " " {
//		return true
//	}
//	tmp := strings.ToLower(s)
//	results := make([]string,0)
//	for i := 0 ;i<len(tmp);i++{
//		if (97 <= uint(tmp[i]) && uint(tmp[i]) <= 122) || (uint(tmp[i]) <= 57 && uint(tmp[i]) >= 48) {
//			results = append(results, string(tmp[i]))
//		}
//	}
//	fmt.Println(results)
//	if len(results) == 0 {
//		return true
//	}
//	for i,j:=0,len(results)-1;len(results) / 2 >= i;i++ {
//		if results[i] == results[j]  {
//			j --
//		} else {
//			return false
//		}
//	}
//	return true
//}

func isPalindrome(s string) bool {

	return true
}

func main() {
	test := "0P"
	fmt.Println(isPalindrome(test))
}
