package main

func isChild(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p || isChild(root.Left, p) || isChild(root.Right, p) {
		return true
	}
	return false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var commonAncestor *TreeNode
	if isChild(root.Left, p) && isChild(root.Left, q) {
		commonAncestor = lowestCommonAncestor(root.Left, p, q)
	} else if isChild(root.Right, p) && isChild(root.Right, q) {
		commonAncestor = lowestCommonAncestor(root.Right, p, q)
	} else {
		commonAncestor = root
	}
	return commonAncestor
}
