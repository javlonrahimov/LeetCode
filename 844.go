package main

func updateString(char *rune, currStr *string) { 
    if (*char == 35 && *currStr != "") {
        *currStr = (*currStr)[:len(*currStr)-1]
    }
    if (*char != 35) {
        *currStr = *currStr+string(*char)
    }
    return
}

func backspaceCompare(S string, T string) bool {
    z,y := "", ""
    for _, char := range S {
        updateString(&char, &z)
    }
    for _, char := range T {
        updateString(&char, &y)
    }
    return y==z
}