package main

import "fmt"

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 打印二叉树
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("空树")
		return
	}

	// 层序遍历打印
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // 弹出队头元素
		fmt.Printf("%d ", node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

// 先序遍历
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	InOrder(root.Right)
}

// 后序遍历
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d ", root.Val)
}
func main() {
	// 测试用例
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	// 打印二叉树
	PrintTree(root)
	fmt.Println("先序遍历:")
	PreOrder(root)
	fmt.Println()
	fmt.Println("中序遍历:")
	InOrder(root)
	fmt.Println()
	fmt.Println("后序遍历:")
	PostOrder(root)
	fmt.Println()
}

/*
二叉树结构图：
       1
      / \
     2   3
    /     \
   4       5
*/
