package main

import "fmt"

// singly-linked list  单链表
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	return
//}

// 数组形式
//func addTwoNumbers(l1 []int, l2 []int) []int {
	resultsl1 := 0
	resultsl2 := 0
	for index,wei:=0,1; index < len(l1); index ++{
		resultsl1 = resultsl1 + (wei * l1[index])
		wei *= 10
	}
	for index,wei:=0,1; index <len(l2); index ++ {
		resultsl2 = resultsl2 + (wei * l2[index])
		wei *= 10
//	}
//	results3 := resultsl1 + resultsl2
//	fmt.Println("results1--->", resultsl1)
//	fmt.Println("results2--->", resultsl2)
//	fmt.Println("results3--->", results3)
//	resultsString := strconv.Itoa(results3)
//	fmt.Println("results3 string --->", resultsString)
//	var r1 []int
//	for i := len(resultsString)-1;i>=0;i--{
//		tmp,_ := strconv.Atoi(string(resultsString[i]))
//		r1 = append(r1, tmp)
//	}
//	return r1
//}
// main 中使用
// list1 := []int{2,4,5}
//	list2 := []int{7,8,8}
//	fmt.Println(addTwoNumbers(list1, list2))

func main() {
	list1 := []int{2, 4, 5}
	head := &ListNode{Val: list1[0]}
	tail := head
	for i := 1; i < len(list1); i++ {
		tail.Next = &ListNode{Val: list1[i]}
		tail = tail.Next
	}
	fmt.Printf("Val: %v,%v", head.Val, head.Next)

}
