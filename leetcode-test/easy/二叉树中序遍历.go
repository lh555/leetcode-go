package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = in(root, res)
	return res
}

func in(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = in(root.Left, res)
	res = append(res, root.Val)
	res = in(root.Right, res)
	return res
}

// 写法2   优先使用这种方式
func inorderTraversalV2(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderTraversalV2(root))
}

/* 对于写法2的解释：
值传递 vs 引用语义：
	Go 语言中，函数参数确实是值传递。当把 ans 传给 dfs（或者在闭包中捕获 ans）时，传递的是这个 SliceHeader 结构体的副本。
	关键点：虽然 Header 是副本，但 Header 里的 Data 指针指向的是同一个底层数组。
append 的行为：
	情况 A（容量足够）：如果 ans 的容量（Cap）足够大，append 直接在底层数组上修改数据，并更新副本中的 Len。由于外层和内层共享底层数组，外层的 ans 能看到数据变化。
	情况 B（容量不足，发生扩容）：如果容量不够，append 会分配一个新的更大的底层数组，把旧数据拷过去，然后返回一个新的 Slice Header（指向新数组）。
	这里有个陷阱：如果在普通函数传参中，func modify(s []int) { s = append(s, 1) }，那么在 modify 内部 s 指向了新数组，但外部的 s 依然指向旧数组，确实会影响不到外层。
	但是！ 你的代码使用的是闭包（Closure）。

在闭包中，ans 并不是作为参数传递进去的副本，而是直接引用了外层函数 inorderTraversalV2 中的那个 ans 变量本身。
当执行 ans = append(...) 时：
append 计算出新的 Slice Header（可能指向新数组）。
赋值操作 ans = ... 直接修改了外层作用域中 ans 变量的值（更新了它的 Data 指针、Len 和 Cap）。
因此，无论是否扩容，外层看到的 ans 永远是最新的。

总结：
1，var dfs func(*TreeNode) 是为了先占位，让后面的匿名函数能通过名字调用自己，实现递归。
2，利用闭包特性，内部函数直接操作外层的 ans 变量。即使 append 导致底层数组扩容并返回新切片，通过 ans = append(...) 赋值，也成功更新了外层变量的引用，所以数据不会丢失。
*/
