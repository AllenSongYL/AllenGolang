package main

import "fmt"

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

// 示例：
//s = "leetcode"
//返回 0

//s = "loveleetcode"
//返回 2

// 哈希表
//func firstUniqChar(s string) int {
//	hashtable := map[string]int{}
//	for i:=0;i<len(s);i++{
//		if _, ok := hashtable[string(s[i])];!ok {
//			hashtable[string(s[i])] = i
//		} else {
//			hashtable[string(s[i])] = -1
//		}
//	}
//	if len(hashtable) == 0 {
//		return -1
//	}
//	fmt.Println(hashtable)
//	temp := []int{-1}
//	for k,v := range hashtable{
//		if hashtable[k] != -1 {
//			if temp[0] == -1 {
//				temp[0] = v
//			} else if temp[0]>v{
//				temp[0] = v
//			} else if temp[0] == 0 {
//				return 0
//			}
//		}
//	}
//	return temp[0]
//}

//  解题
func firstUniqChar(s string) int {
	var arr [26]int
	// 先进行第一次遍历，在数组中记录每个字母的最后一次出现的所在索引
	for i, k := range s {
		arr[k-'a'] = i
		//fmt.Println(arr[k - 'a'])
	}
	// 然后再通过一次循环，比较各个字母第一次出现的索引是否为最后一次的索引。
	//如果是，我们就找到了我们的目标，
	//如果不是我们将其设为 -1（标示该元素非目标元素）
	//如果第二次遍历最终没有找到目标，直接返回 -1即可。
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		} else {
			arr[k-'a'] = -1
		}
	}
	return -1
}

func main() {
	test := "cc"
	fmt.Println(firstUniqChar(test))
	//firstUniqChar(test)
}
