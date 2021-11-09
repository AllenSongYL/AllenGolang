package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	fast := head
	//slow := &ListNode{}
	var slow *ListNode
	//tmp2 := results
	//nil--->10--->20--->30--->40--->nil
	//   slow fast  tmp
	//           slow fast  tmp
	for fast != nil {
		if fast.Next == nil {
			fast.Next = slow
			break
		}
		tmp := fast.Next
		fast.Next, slow = slow, fast
		fast = tmp
	}
	return fast
}

func main() {
	head := &ListNode{Val: 10}
	h1 := &ListNode{Val: 20}
	head.Next = h1
	h2 := &ListNode{Val: 30}
	h1.Next = h2
	h3 := &ListNode{Val: 40}
	h2.Next = h3
	t := reverseList(head)
	t2 := t
	for {
		fmt.Println(t2.Val)
		if t2.Next == nil {
			break
		}
		t2 = t2.Next
	}
}
