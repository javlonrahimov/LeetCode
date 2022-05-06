package main

import "strings"

func removeDuplicates(s string, k int) string {
        var sb strings.Builder
    stack := []rune{}
    count := []int{}
    for _, char := range s{
        if len(stack) == 0{
            stack = append(stack, char)
            count = append(count, 1)
        }else{
            if stack[len(stack)-1] != char{
                stack = append(stack, char)
                count = append(count, 1)
            }else{
                if count[len(count)-1] == k-1{
                    count= count[:len(count)-1]
                    stack = stack[:len(stack)-1]
                }else{
                    count[len(count)-1]++
                }
            }
        }
    }
    for i := range stack {        
    	for j:=0; j<count[i]; j++{
            sb.WriteString(string(stack[i]))
        }
    }
    ret := sb.String()
    return ret
}
