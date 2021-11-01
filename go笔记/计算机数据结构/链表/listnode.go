package main

import "fmt"

type HeroNode struct {
	No       int
	Name     string
	NickName string
	Next     *HeroNode
}

// 第一种插入方式：在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到该链表的最后节点
	temp := head
	for {
		// 表示找到最后节点
		if temp.Next == nil {
			break
		}
		// 让temp不断指向下一个节点
		temp = temp.Next
	}
	// 将newHeroNode 加入到链表的最后
	temp.Next = newHeroNode
}

// InsertHeroNode2 第二种方式：根据id大小，从小到大进行插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	// 让插入节点的id和temp.Next的下一个id进行比较
	flag := true
	for {
		if temp.Next == nil {
			break
		} else if newHeroNode.No < temp.Next.No {
			// 添加的比 temp后面的id小，插入到temp后面
			//newHeroNode.Next = temp.Next
			//temp.Next = newHeroNode
			break
		} else if temp.Next.No == newHeroNode.No {
			flag = false
			fmt.Println("存在重复ID")
			break
		}
		temp = temp.Next
	}
	if flag == false {
		fmt.Println("已存在")
		return
	} else {
		newHeroNode.Next = temp.Next
		temp.Next = newHeroNode
	}
}

// ListSingleList 显示链表的所有节点信息
func ListSingleList(head *HeroNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("空链表！")
		return
	}
	for {
		fmt.Printf("节点信息：ID: %d, Name: %s, NickName: %v\n",
			temp.Next.No, temp.Next.Name, temp.Next.NickName)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
}

func main() {
	// 创建头节点，可以没有值
	head := &HeroNode{}

	// 创建一个新的HeroNode
	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}
	hero2 := &HeroNode{
		No:       2,
		Name:     "卢俊义",
		NickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		No:       3,
		Name:     "林冲",
		NickName: "豹子头",
	}
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero1)
	ListSingleList(head)

}
