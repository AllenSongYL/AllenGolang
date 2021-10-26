package main

import (
	"fmt"
)

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 位运算
// 1^1 = 0 自己和自己异或等于0
// a^0 = a 任何数字和0异或还等于他自身
// a^b^c=a^c^b  a^b^a = a^a^b = b 异或运算具有交换律
func singleNumber(nums []int) int {
	var results int
	for _, v := range nums {
		results = results ^ v
	}
	return results
}

func main() {
	list1 := []int{4, 2, 1, 2, 4}
	fmt.Println(singleNumber(list1))

}
