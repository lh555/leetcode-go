package main

import "fmt"

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

class Solution(object):

	def isPalindrome(self, head):
	    """
	    :type head: Optional[ListNode]
	    :rtype: bool
	    """
	    if not head:
	        return False
	    cache = []
	    cur = head
	    while cur:
	        cache.append(cur.val)
	        cur = cur.next
	    return cache == cache[::-1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	cache := make([]int, 0)
	cur := head
	for cur != nil {
		cache = append(cache, cur.Val)
		cur = cur.Next
	}
	// 反转切片比较
	n := len(cache)
	for i, v := range cache[0 : n/2] {
		if v != cache[n-1-i] {
			return false
		}
	}
	return true
}

// 快慢指针 + 反转后半链表，空间 O(1)
func isPalindromeV2(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 1. 快慢指针找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半链表
	var prev *ListNode
	for slow != nil {
		next := slow.Next // 先保存下一个节点
		slow.Next = prev  // 反转当前节点指向
		prev = slow       // prev 前移
		slow = next       // slow 后移（用保存的 next）
	}
	// prev 现在是反转后的头

	// 3. 前后对比
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func main() {
	list1 := &ListNode{1, nil}
	list1.Next = &ListNode{2, nil}
	list1.Next.Next = &ListNode{2, nil}
	list1.Next.Next.Next = &ListNode{1, nil}
	res := isPalindromeV2(list1)
	fmt.Println(res)
}
