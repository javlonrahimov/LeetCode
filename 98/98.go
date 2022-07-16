package main

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil &&
		(root.Val <= root.Left.Val ||
			(min != nil && min.Val >= root.Left.Val)) {
		return false
	}

	if root.Right != nil &&
		(root.Val >= root.Right.Val ||
			(max != nil && max.Val <= root.Right.Val)) {
		return false
	}

	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}
