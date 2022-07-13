package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	l := [][]int{
		[]int{root.Val},
	}

	for s, cc := helper([]*TreeNode{root.Left, root.Right}); s != nil; s, cc = helper(cc) {
		l = append(l, s)
	}

	return l
}

func helper(c []*TreeNode) ([]int, []*TreeNode) {
	if c == nil {
		return nil, nil
	}
	var (
		s  []int
		cc []*TreeNode
	)

	for _, n := range c {
		if n == nil {
			continue
		}
		s = append(s, n.Val)
		cc = append(cc, []*TreeNode{n.Left, n.Right}...)
	}
	return s, cc
}
