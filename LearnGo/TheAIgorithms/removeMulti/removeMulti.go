package removeMulti

import "fmt"

// 传入已排序slice
func RemoveMultiNums(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}

	var left int = 0
	var right int = 1
	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
			//fmt.Println(nums[i])
			//fmt.Println(nums[i-1])
		}
		left++
		nums[left] = nums[right]
		fmt.Println("循环体中: ", nums)
	}
	fmt.Println(nums[:left+1])
	return 0
}
