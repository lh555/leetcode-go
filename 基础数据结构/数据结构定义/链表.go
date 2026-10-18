package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// PrintList 打印链表内容
func PrintList(head *ListNode) {
	current := head
	if current == nil {
		fmt.Println("空链表")
		return
	}

	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// CreateListFromSlice 从切片创建链表
func CreateListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

func main() {
	fmt.Println("请输入一串数字，用空格分隔:")

	// 读取用户输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// 去除换行符并分割字符串
	input = strings.TrimSpace(input)
	strNums := strings.Split(input, " ")
	fmt.Println("输入的数字为:", strNums)
	// 转换为整数切片
	var nums []int
	for _, str := range strNums {
		if num, err := strconv.Atoi(str); err == nil {
			nums = append(nums, num)
		}
	}

	if len(nums) == 0 {
		fmt.Println("未输入有效数字")
		return
	}

	// 从输入创建链表
	head := CreateListFromSlice(nums)

	fmt.Println("原始链表:")
	PrintList(head)

	// 反转链表
	reversed := ReverseList(head)

	fmt.Println("反转后链表:")
	PrintList(reversed)
}
