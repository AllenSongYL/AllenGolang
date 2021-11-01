package main

import (
	"fmt"
)

//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	nums = nums[:slow+1]
	fmt.Println(nums)
	return slow + 1
}

func main() {
	test := []int{1, 2, 3, 4, 5, 5, 6, 6, 6}
	removeDuplicates(test)
}
