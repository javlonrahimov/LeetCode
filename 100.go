package main

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {return true};

	if q == nil || p == nil {return false};

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
