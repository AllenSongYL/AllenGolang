package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			j++
		}
	}
	return j
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 7, 7}
	removeElement(test, 7)
	fmt.Println(test)
}
