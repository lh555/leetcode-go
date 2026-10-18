package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	list1 := &ListNode{1, nil}
	list1.Next = &ListNode{2, nil}
	list1.Next.Next = &ListNode{3, nil}
	res := reverseList(list1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
