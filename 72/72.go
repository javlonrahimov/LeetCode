package main

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func solve(i, j int, word1, word2 string, dp [][]int) int {
    if i == len(word1) && j == len(word2) {
        return 0
    }else if i == len(word1){ 
        return len(word2) - j
    }else if j == len(word2) {
        return len(word1) - i
    }
    if dp[i][j] != -1 {
        return dp[i][j]
    }
    if word1[i] == word2[j] {
        dp[i][j] = solve(i + 1, j + 1, word1, word2, dp)
        return dp[i][j]
    }
    dp[i][j] = solve(i, j + 1, word1, word2, dp)
    dp[i][j] = min(dp[i][j], solve(i + 1, j, word1, word2, dp))
    dp[i][j] = min(dp[i][j], solve(i + 1, j + 1, word1, word2, dp))
    dp[i][j]++
    return dp[i][j]
}

func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1))
    for i := 0; i < len(word1); i++ {
        dp[i] = make([]int, len(word2))
        for j := 0; j < len(word2); j++ {
            dp[i][j] = -1
        }
    }
    return solve(0, 0, word1, word2, dp)
}
