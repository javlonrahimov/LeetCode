package main

func findSubstring(s string, words []string) []int {
    wordLength := len(words[0])
    wordsLength := len(words)
    windowLength := wordLength * wordsLength
    wordsCount := make(map[string]int)
    wordsCountSum := 0
    for i := 0; i < len(words); i++ {
        wordsCount[words[i]]++
        wordsCountSum++
    }
    var result []int
    for i := 0; i <= len(s) - windowLength; i++ {
        if _, ok := wordsCount[s[i:i+wordLength]]; !ok {
            continue
        }
        windowCompareCount := copyMap(wordsCount)
        windowCompareCount[s[i:i+wordLength]]--
        countSum := wordsCountSum - 1
        for j := 1; j < wordsLength; j++ {
            count, ok := windowCompareCount[s[i + j * wordLength: i + j * wordLength + wordLength]]
            if !ok || count == 0 || wordsCountSum < 0 {
                break
            }
            windowCompareCount[s[i + j * wordLength: i + j * wordLength + wordLength]]--
            countSum--
        }
        if countSum == 0 {
            result = append(result, i)
        }
    }
    return result
}

func copyMap(original map[string]int) map[string]int {
    result := make(map[string]int)
    for key, value := range original {
        result[key] = value
    }
    return result
}
