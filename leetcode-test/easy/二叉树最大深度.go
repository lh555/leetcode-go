package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDeep := maxDepth(root.Left)
	rightDeep := maxDepth(root.Right)
	return max(leftDeep, rightDeep) + 1
}
