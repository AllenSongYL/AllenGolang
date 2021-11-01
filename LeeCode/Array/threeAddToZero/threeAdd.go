package main

import (
	"fmt"
)

// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

// 尝试三指针
//func threeSum(nums []int) {
//	if len(nums) < 3 {
//	}
//	results := make([][]int, len(nums)/3)
//	for i,_ := range results {
//		results[i] = make([]int, 3)
//	}
//	for fast,middle,slow,index,counts := 2,1,0,0,3;fast <= len(nums);fast ++ {
//		fmt.Println(fast)
//		if nums[fast] + nums[middle] + nums[slow] == 0 {
//			//resutlts = append(results, [][]int{nums[fast], nums[middle], nums[slow]})
//			results[index][0] = nums[slow]
//			results[index][1] = nums[middle]
//			results[index][2] = nums[fast]
//			if len(nums)-counts < 3 {
//
//			}
//			index ++
//			fmt.Println("results", results)
//		} else if fast == len(nums) {
//			middle ++
//			fast = middle +1
//		}
//	}
//	//return results
//}

func main() {
	test := []int{-1, 1, 0, 2, -2}
	threeSum(test)
}
