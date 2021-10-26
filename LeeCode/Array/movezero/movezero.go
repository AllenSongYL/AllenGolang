package main

import "fmt"

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
// 输入: [0,1,0,3,12,0]
// 输出: [1,3,12,0,0]

func moveZeroes(nums []int) []int {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zero], nums[i] = nums[i], nums[zero]
			zero++
		}

	}
	return nums
}

func main() {
	sliceA := []int{0, 1, 4, 0, 8, 9, 0, 0, 33}
	fmt.Println(moveZeroes(sliceA))
}
