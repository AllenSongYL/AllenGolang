package main

import "fmt"

func lengthOfLastWord(s string) int {
	var results int
	for n := len(s) - 1; n >= 0; n-- {
		if string(s[n]) == " " {
			continue
		} else if string(s[n]) != " " {
			results++
			if string(s[n-1]) == " " {
				return results
			}
		}
	}
	return results
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
