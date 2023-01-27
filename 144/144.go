package main


func preorderTraversal(root *TreeNode) []int {
	if root == nil {
        return nil
    }

    var ret []int
    p := root
    stack := make([]*TreeNode, 0)
    stack = append(stack, p)

    for len(stack) != 0 {
        visit := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        ret = append(ret, visit.Val)
        
        if visit.Right != nil {
            stack = append(stack, visit.Right)
        }
        if visit.Left != nil {
            stack = append(stack, visit.Left)
        }
    }
    return ret    
}
