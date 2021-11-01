package main

import (
	"fmt"
	"strings"
)
import (
	path2 "path"
)

func simplifyPath(path string) string {
	return path2.Clean(path)
}

//func simplifyPath(path string) string {
//	result1 := strings.Replace(path, "/./","",-1)
//	result2 := strings.Replace(result1, "//","/",-1)
//	fmt.Println(result2)
//	result := []string{"/"}
//	results := ""
//	tmp := 0
//	for i:=1;i<len(result2);i++ {
//		//fmt.Println(len(result2))
//		if string(result2[i]) == "/" || i ==len(result2)-1 {
//			result = append(result, string(result2[tmp:i+1]))
//			tmp = 0
//			continue
//			//&& (string(result2[i+2]) == "/" || i+1 == len(result2)-1 )
//		} else if string(result2[i]) == "." && string(result2[i+1]) == "." && string(result2[i+2])!="." {
//			if len(result) == 1 {
//				i = i + 2
//			} else {
//				result = result[0:len(result)-1]
//				i = i + 2
//			}
//		}else {
//			if tmp == 0 {
//				tmp = i
//			}
//		}
//	}
//	//return string(result[0])
//	for _,value := range result{
//		results = results + value
//	}
//	if string(results[len(results)-1]) == "/" && len(result) != 1{
//		results =  results[0:len(results)-1]
//	}
//	return  results
//}

func simplifyPath(path string) string {
	a := strings.Split(path, "/")
	results := ""
	tmp := []string{""}

	for _, v := range a {
		if v != "" && v != "." {
			if v == ".." {
				if len(tmp)-1 <= 0 {
					tmp[0] = ""
				} else {
					tmp = tmp[:len(tmp)-1]
				}
			} else {
				v = "/" + v
				tmp = append(tmp, v)
			}
		}
	}
	for _, value := range tmp {
		results = results + value
	}
	if len(results) == 0 {
		results = "/"
	}
	fmt.Println(results)
	return results
}

func main() {

	//test := "/a/./b/../../c/"
	//test := "/home/aaa/ccc/../dd/"
	//test:="/../"
	//test := "/a//b////c/d//././/.."
	test := "/a/../../b/../c//.//"
	fmt.Println(simplifyPath(test))
}
