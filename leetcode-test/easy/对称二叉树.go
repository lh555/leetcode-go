package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return recur(root.Left, root.Right)
}

func recur(left *TreeNode, right *TreeNode) bool {
	// 都为nil则对称
	if left == nil && right == nil {
		return true
	}
	// 其中一个为nil则不对称
	if left == nil || right == nil {
		return false
	}
	// 值不相等则不对称
	if left.Val != right.Val {
		return false
	}
	return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}
