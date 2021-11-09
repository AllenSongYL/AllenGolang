package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if n <=0 {
//		return head
//	}
//	tmp := head
//	//counts := 1
//	for {
//		// 调用删除为1
//		if n == 1 {
//			// 如果链表只有一个head && counts == 1
//			if tmp.Next == nil {
//				head.Val = 0
//				head.Next = nil
//			// 如果链表长度大于1 && counts >1
//			} else if tmp.Next.Next == nil  {
//				tmp.Next = nil
//			}
//			return head
//		}
//
//		// 循环结束条件
//		if tmp.Next == nil {
//			break
//		}
//
//		// tmp2 保存快指针
//		//tmpfast := tmp.Next
//
//		// 10,20,30,40,50  4
//		//for i:=1; i<n && tmpfast.Next != nil; i++ {
//		//	tmpfast = tmpfast.Next.Next
//		//}
//		//
//		//if tmpfast.Next == nil && tmp == head{
//		//	tmp3 := head.Next
//		//	head.Next = nil
//		//	fmt.Println("tmp3")
//		//	return tmp3
//		//} else if tmpfast.Next == nil {
//		//	tmp.Next = tmp.Next.Next
//		//	return head
//		//}
//		tmp = tmp.Next
//		//counts ++
//	}
// 	return head
//}

// good way
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	mynewHead := &ListNode{}
//	// 把传入的head 放到mynewHead下
//	mynewHead.Next = head
//	// 删除的结点
//	beforDel := head
//	// 上结点
//	prev := mynewHead
//	i := 1
//	// 0 10，20，30，40，50 5
//	for beforDel != nil {
//		beforDel = beforDel.Next
//		if i > n {
//			prev = prev.Next
//		}
//		i++
//	}
//	prev.Next = prev.Next.Next
//	return mynewHead.Next
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newhead := &ListNode{}
	newhead.Next = head
	delbefore := head
	pre := newhead
	i := 1
	//0 1,2,3,4,5,6  1
	for delbefore != nil {
		delbefore = delbefore.Next
		if i > n {
			pre = pre.Next
		}
		i++
	}
	pre.Next = pre.Next.Next
	return newhead.Next
}

func main() {
	head := &ListNode{Val: 10}
	h1 := &ListNode{Val: 20}
	head.Next = h1
	h2 := &ListNode{Val: 30}
	h1.Next = h2
	h3 := &ListNode{Val: 40}
	h2.Next = h3
	h4 := &ListNode{Val: 50}
	h3.Next = h4
	//

	tmp := removeNthFromEnd(head, 1)
	for i := 0; ; i++ {
		fmt.Println(tmp.Val)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
}
