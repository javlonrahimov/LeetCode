package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 1 + Max(MaxDepth(root.Left), MaxDepth(root.Right))
	return count

}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
