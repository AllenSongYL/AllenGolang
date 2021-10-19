package main

import (
	"fmt"
	"go_env/LearnGo/TheAIgorithms/removeMulti"
)

func main() {

	nums := make([]int, 10)
	nums[0] = 1
	nums[1] = 1
	nums[2] = 2
	nums[3] = 2
	nums[4] = 4
	nums[5] = 7
	nums[6] = 8
	nums[7] = 9
	nums[8] = 9
	nums[9] = 9
	fmt.Println(nums)
	removeMulti.RemoveMultiNums(nums)
	//fmt.Println(nums)
}
