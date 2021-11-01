package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
//如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:
//输入: nums = [1,3,5,6], target = 5
//输出: 2

//示例2:
//输入: nums = [1,3,5,6], target = 2
//输出: 1

//示例 3:
//输入: nums = [1,3,5,6], target = 7
//输出: 4

//示例 4:
//输入: nums = [1,3,5,6], target = 0
//输出: 0

//示例 5:
//输入: nums = [1], target = 0
//输出: 0

// 题解 大于等于目标值则直接返回索引位
//  否则放到数组的最后
func searchInsert(nums []int, target int) int {
	for key, value := range nums {
		if value >= target {
			return key
		}
	}
	return len(nums)
}

func main() {
	test := []int{1, 3, 6}
	fmt.Println(searchInsert(test, 5))
}
