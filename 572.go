package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
    
    if s==nil{
        return false
    }
    
    if s.Val == t.Val && isEqual(s, t){
        return true
    }else{
        return isSubtree(s.Left, t) || isSubtree(s.Right, t)
    }
}


func isEqual(s *TreeNode, t *TreeNode) bool{
    
    if s==nil&& t==nil{
        return true
    }
    
    if s==nil|| t==nil{
        return false
    }
    
    if s.Val == t.Val{
        return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
    }else{
        return false
    }
}