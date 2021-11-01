package main

import "fmt"

func isSelfCrossing(distance []int) bool {
	if len(distance) <= 3 {
		return false
	}
	resultsUp := distance[0] - distance[2]
	resultsLeft := distance[1] - distance[3]
	if len(distance) == 4 {
		if resultsUp >= 0 && resultsLeft <= 0 {
			return true
		}
		return false
	} else {
		if resultsUp >= 0 && resultsLeft <= 0 {
			return true
		} else {
			for one := 4; one < len(distance); one++ {
				if resultsUp := resultsUp + distance[one]; resultsUp <= 0 {
					return false
				}
			}
		}
	}
	return false
}

func main() {
	test := []int{1, 1, 2, 2, 1, 1}
	fmt.Println(isSelfCrossing(test))
}
