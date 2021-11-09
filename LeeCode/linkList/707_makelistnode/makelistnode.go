package main

import (
	"fmt"
)

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

// Get方法: 通过输入索引位，返回该位置的值
func (this *MyLinkedList) Get(index int) int {
	tmp := this
	if this.Next == nil {
		return -1
	}
	counts := 0
	for {
		tmp = tmp.Next
		if counts == index {
			return tmp.Val
		}
		counts++
		if tmp.Next == nil {
			return -1
		}
	}
}

// AddAtHead方法：在链表的第一个元素之前添加一个值为 val 的节点。
// 插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	// 新建一个newhead，Next指向原本的第一位
	newhead := &MyLinkedList{Val: val}
	newhead.Next = this.Next
	// 头结点指向新建的newhead
	this.Next = newhead
}

// 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := this
	new2 := &MyLinkedList{Val: val}
	for {
		// 表示找到最后节点
		if tmp.Next == nil {
			tmp.Next = new2
			break
		}
		// 让temp不断指向下一个节点
		tmp = tmp.Next
	}
}

// 在链表中的第index个节点之前添加值为val的节点。如果index等于链表的长度，则该节点将附加到链表的末尾。
// 如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// index小于0 则在头部插入
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	new3 := &MyLinkedList{Val: val}
	tmp := this
	counts := 0

	for {
		// 循环到指定索引，新结点指向原始后面结点，
		if counts == index {
			new3.Next = tmp.Next
			tmp.Next = new3
			break
		}
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
		counts++
	}
}

// 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	tmp := this
	counts := 0
	for {
		// 删除最后一个
		//if tmp.Next.Next == nil && counts+1 == index{
		//	tmp.Next = nil
		//	break
		//}

		// 超过索引 则退出

		// && tmp.Next.Next != nil
		if counts == index && tmp.Next != nil {
			tmp.Next = tmp.Next.Next
			break
		}
		counts++
		if counts > index || tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
}

func main() {
	// 新建头节点
	head := new(MyLinkedList)
	// 第一个结点
	//h1 := &MyLinkedList{Val: 10}
	//head.Next = h1
	// 第二个结点
	//h2 := &MyLinkedList{Val: 20}
	//h1.Next = h2
	//h3 := &MyLinkedList{Val: 30}
	//h2.Next = h3
	fmt.Println("添加")
	head.AddAtHead(2)
	head.AddAtIndex(0, 1)
	fmt.Println(head.Get(3))
	//head.AddAtHead(33)

	//head.AddAtTail(44)
	//head.AddAtTail(55)
	//fmt.Println(head.Get(0))

	// 33,10,20,44,55
	//head.AddAtIndex(5,700)
	// 33,10,20,44,55,700
	//head.DeleteAtIndex(8)

	//tmp := head.Next
	//for i := 0;;i++ {
	//	if tmp.Next == nil {
	//		break
	//	}
	//	tmp = tmp.Next
	//}

}
