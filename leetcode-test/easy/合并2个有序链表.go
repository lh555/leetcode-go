package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// 迭代
func mergeTwoListsV2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	list3 := &ListNode{0, nil}
	p1, p2, p3 := list1, list2, list3
	for p1 != nil && p2 != nil {
		tmp := &ListNode{0, nil}
		if p1.Val < p2.Val {
			tmp.Val = p1.Val
			p1 = p1.Next
			p3.Next = tmp
			p3 = p3.Next
		} else {
			tmp.Val = p2.Val
			p2 = p2.Next
			p3.Next = tmp
			p3 = p3.Next
		}
	}
	if p1 == nil {
		p3.Next = p2
	} else {
		p3.Next = p1
	}
	return list3.Next
}

func main() {
	list1 := &ListNode{1, nil}
	list1.Next = &ListNode{2, nil}
	list1.Next.Next = &ListNode{4, nil}
	list2 := &ListNode{1, nil}
	list2.Next = &ListNode{3, nil}
	list2.Next.Next = &ListNode{4, nil}

	// 递归
	recur := mergeTwoLists(list1, list2)
	for ; recur != nil; recur = recur.Next {
		fmt.Print(recur.Val)
	}
	fmt.Println()
	// 迭代
	list1 = &ListNode{1, nil}
	list1.Next = &ListNode{2, nil}
	list1.Next.Next = &ListNode{4, nil}
	list2 = &ListNode{1, nil}
	list2.Next = &ListNode{3, nil}
	list2.Next.Next = &ListNode{4, nil}
	res := mergeTwoListsV2(list1, list2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}
