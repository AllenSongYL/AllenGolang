package main

import "fmt"

// 给定一个整数数组，判断是否存在重复元素。
//如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

// 输入: [1,2,3,1]          [1,2,3,4]
// 输出: true                false

func containsDuplicate(nums []int) bool {
	hashtable := make(map[int]int)
	for _, value := range nums {
		if hashtable[value] != 0 {
			return true
		} else {
			hashtable[value] = 1
		}
	}
	return false
}

func main() {
	amap := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(amap))

}
