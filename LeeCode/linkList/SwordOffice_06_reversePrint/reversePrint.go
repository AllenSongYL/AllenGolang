package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var results []int
	tmp := head

	if tmp == nil {
		return results
	}

	for {
		results = append(results, tmp.Val)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	return results
}

func main() {

	head := &ListNode{Val: 1}
	h1 := &ListNode{Val: 10}
	head.Next = h1
	h2 := &ListNode{Val: 20}
	h1.Next = h2
	h3 := &ListNode{Val: 30}
	h2.Next = h3
	fmt.Println(reversePrint(head))

}
