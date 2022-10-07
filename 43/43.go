package main

import "strconv"


func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    total := make([]int, len(num1) + len(num2))
    
    for i, d1 := range num1 {
        for j, d2 := range num2 {
            tensPlace := (len(num1) - i - 1) + (len(num2) - j - 1)
            value := int(d1 - '0') * int(d2 - '0')
            
            total[tensPlace] += value
        }
    }
    
    res := ""
    
    for i, num := range total {
        if i < len(total) - 1 {
            total[i + 1] += num / 10
        }
        
        res = strconv.Itoa(num % 10) + res
    }
    
    if res[0] == '0' {
        res = res[1:]
    }
    
    return res
}