package main

import "fmt"

// 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。
// 由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
// 你可以通过调用bool isBadVersion(version)接口来判断版本号 version 是否在单元测试中出错。
// 实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

// 示例 1：
// 输入：n = 5, bad = 4
// 输出：4

//解释：
//调用 isBadVersion(3) -> false
//调用 isBadVersion(5)-> true
//调用 isBadVersion(4)-> true
//所以，4 是第一个错误的版本。

//示例 2：
//输入：n = 1, bad = 1
//输出：1

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var bad = 4

func isBadVersion(versions int) bool {
	if versions >= bad {
		return true
	}
	return false
}

// stack overflow
//func firstBadVersion(n int) int {
//	if isBadVersion(n) == true && isBadVersion(n+1) == false {
//		return n+1
//	} else if isBadVersion(n) == true && isBadVersion(n+1) == true {
//		firstBadVersion(n+2)
//	} else if isBadVersion(n) == false && isBadVersion(n-1) == true {
//		return n
//	} else if isBadVersion(n) == false && isBadVersion(n-1) == false {
//		firstBadVersion(n-2)
//	}
//	return 0
//}

func firstBadVersion(n int) int {
	// 输入的是坏版本，且前一个是正确的
	if isBadVersion(n) == true && isBadVersion(n-1) == false {
		return n
		// 输入的是正确的，且后一个是错误的
	} else if isBadVersion(n) == false && isBadVersion(n+1) == true {
		return n + 1
	} else if isBadVersion(n) == true {
		for i := n - 1; n >= 0; i-- {
			if isBadVersion(i-1) == false {
				return i
			}
		}
	} else if isBadVersion(n) == false {
		for i := n + 1; n >= 0; i++ {
			if isBadVersion(i+1) == true {
				return i + 1
			}
		}

	}
	return -1
}

func main() {
	//test := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(firstBadVersion(5))

}
