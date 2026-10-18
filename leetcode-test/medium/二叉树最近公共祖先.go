package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如果一个节点是 p 和 q 的最近公共祖先（LCA），那么 p 和 q 应该分别位于该节点的左子树和右子树中，或者该节点本身就是 p 或 q。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 当前层的左右都不为空，说明p，q在当前节点（root）左右，那当前节点就是LCA。返回root
	if left != nil && right != nil {
		return root
	}
	// 左子树为空，说明p，q都在右子树，那右子树的返回值就是LCA。返回right
	// 右子树为空，说明p，q都在左子树，那左子树的返回值就是LCA。返回left
	if left != nil {
		return left
	} else {
		return right
	}
}
