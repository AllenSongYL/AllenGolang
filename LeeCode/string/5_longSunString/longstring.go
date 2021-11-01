package main

import (
	"fmt"
)

//func longestPalindrome(s string) string {
//	if len(s) < 2 {
//		return string(s[0])
//	}
//    if len(s) == 2 {
//        if s[0] == s[1] {
//            return s
//        } else {
//            return string(s[0])
//        }
//    }
//
//    results := map[string]int{}
//
//	for i,l:=0,len(s)-1;l>=0;l-- {
//		if s[i] == s[l] && l>i{
//			results[s[i:l+1]]=l-i+1
//            if i<len(s)-1{
//                i ++
//                l = len(s) - 1
//            }
//
//
//		} else if l>i {
//			l --
//		}
//		if l == i{
//			l = len(s) - 1
//			i ++
//		}
//	}
//
//	for str,lens := range results {
//		temp := 0
//		if lens > temp {
//			temp = lens
//		}
//		return str
//	}
//	return results
//}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return string(s[0])
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	temps := make([]string, 1)
	if len(s)%2 != 0 {
		for i := 1; i < len(s)-1; i++ {
			for left, right := i-1, i+1; left >= 0 && right < len(s); {
				if s[left] == s[right] {
					if len(temps[0]) < right-left+1 {
						temps[0] = s[left : right+1]
					}
				} else if s[i] == s[left] {
					if len(temps[0]) < 2 {
						temps[0] = s[left : i+1]
					}
				} else if s[i] == s[right] {
					if len(temps[0]) < 2 {
						temps[0] = s[i : right+1]
					}
				}
				left--
				right++
			}
		}
	} else {
		for start1, start2 := 0, 1; start2 < len(s); {
			if s[start1] == s[start2] {
				if len(temps[0]) < start2-start1 {
					temps[0] = s[start1 : start2+1]
				}
				if start1-1 > 0 && start2+1 < len(s) {
					start1--
					start2++
				}
			}
			start1++
			start2++

		}
		return temps[0]
	}
	fmt.Println(temps)

	return temps[0]
}

func main() {
	//test := "abb"
	//test := "ccd"
	test := "aaaa"
	fmt.Println(longestPalindrome(test))

}
