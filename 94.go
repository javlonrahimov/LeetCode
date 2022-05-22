package main


type stack []TreeNode

func (s stack) Push(v TreeNode) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, *TreeNode) {
    // FIXME: What do we do if the stack is empty, though?

    l := len(s)
    if(l == 0) {
        return nil, nil
    }
    return  s[:l-1], &s[l-1]
}

func inorderTraversal(root *TreeNode) []int {
 	   res := make([]int, 0)
 	   stack := make(stack, 0)
 	   curr := root
 	   for curr != nil || (len(stack) != 0) {
 	   	for curr != nil {
 	   		stack.Push(*curr)
 	   		curr = curr.Left
 	   	}
 	   	 _, curr = stack.Pop()
 	   	res = append(res, curr.Val)
 	   	curr = curr.Right
 	   }
 	   return res
}
//todo complete
