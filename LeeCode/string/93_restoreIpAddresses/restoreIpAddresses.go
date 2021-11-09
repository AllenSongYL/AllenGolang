package main

import (
	"fmt"
	"unicode"
)

//func restoreIpAddresses(s string) []string {
//	var results []string
//	if len(s) < 4 {
//		return []string{}
//	}
//
//	if len(s) == 4 {
//		var resstr string
//		for _,value := range s {
//			resstr = resstr + string(value) + "."
//		}
//		resstr = resstr[:len(resstr)-1]
//		results = append(results, resstr)
//		return results
//	}
//
//	for i,j:=0,2;i<len(s);i++{
//		var tmp string
//		if s[0] == 0 {
//            tmp = "0."
//            j++
//            break
//        }
//
//	}
//	return results
//}

func main() {
	test := "2113"
	fmt.Println(restoreIpAddresses(test))
}
