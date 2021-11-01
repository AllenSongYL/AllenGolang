package main

import (
	"fmt"
)

//func longestCommonPrefix(strs []string) string {
//	results := make([]string,0)
//	for _,value := range strs {
//		for i,j := range value {
//			if results[i] == string(j){
//				continue
//			} else {
//				results[i] = string(j)
//			}
//		}
//	}
//	if len(results) == 0{
//		return ""
//	}
//	var resu string
//	for i:=0;i<len(results);i++{
//		resu =  resu + string(resu[i])
//	}
//	return resu
//}

func longestCommonPrefix(strs []string) string {
	// "flower"
	if len(strs) == 1 {
		return strs[0]
	}

	results := strs[0]

	if len(results) == 0 {
		return ""
	}
	var index int
	// "flow","flight"
	for _, value := range strs[1:] {
		if len(value) == 0 {
			return ""
		}
		for i := 1; i <= len(value) && i <= len(results); i++ {
			if results[0] != value[0] {
				return ""
			}
			if value[:i] == results[:i] {
				index = i + 1
				continue
			} else {
				index = i
				break
			}
		}
		if index == 0 {
			results = string(results[0])
		} else {
			results = results[:index-1]
		}
	}
	return results
}

func main() {

	test := []string{
		//"flower","flow","flight",
		//"dog","racecar","car",
		"ab", "a",
		//"","",
		//"flower","flower","flower","flower",
		//"flower","fkow",
		//"a","a","b",
	}
	fmt.Println(longestCommonPrefix(test))
}
