package main

func isSymmetric(root *TreeNode) bool {
	return helper101(root, root)
}

func helper101(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && helper101(left.Right, right.Left) && helper101(left.Left, right.Right)

}
