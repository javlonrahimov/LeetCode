package main

func rightSideView(root *TreeNode) []int {
 	res := []int{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        newLevel := []int{}
        for i := 0; i < n; i++ {
            node := queue[0]
            queue = queue[1:]
            newLevel = append(newLevel, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            } 
        }
		// The right-side view is simply the last slice element at each level.
        res = append(res, newLevel[len(newLevel)-1])
    }
    return res
}