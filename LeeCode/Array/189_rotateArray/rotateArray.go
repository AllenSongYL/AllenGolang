package main

import "fmt"

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
// 输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]

// 输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]

func rotate(nums []int, k int) {
	if len(nums) < 2 || k < 1 {
		fmt.Println("nums", nums)
		return
	}
	temp := k - len(nums)
	if k > len(nums) {
		if k%len(nums) == 0 {
			fmt.Println("nums", nums)
			return
		}
		temp = len(nums) - (k % len(nums))
		//fmt.Println("商:", temp)
	} else if k == len(nums) {
		fmt.Println("nums", nums)
		return
	} else {
		temp = -temp
	}
	nums1 := nums[temp:]
	fmt.Println("nums1: ", nums1)
	nums2 := nums[:temp]
	fmt.Println("nums2: ", nums2)
	nums = append(nums[0:0], append(nums1, nums2...)...)
	fmt.Println("nums", nums)
}

func main() {
	test := []int{-1, 2, 3}
	rotate(test, 7)
}
