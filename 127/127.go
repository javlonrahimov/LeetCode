package main

type node struct {
	word  string
	state int
}

func LadderLength(beginWord, endWord string, wordList []string) int {

	l := len(beginWord)

	dict := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < l; j++ {
			dict[wordList[i][:j]+"*"+wordList[i][j+1:]] = append(dict[wordList[i][:j]+"*"+wordList[i][j+1:]], wordList[i])
		}
	}

	queue := []node{{word: beginWord, state: 1}}
	set := make(map[string]struct{})
	set[beginWord] = struct{}{}
	for len(queue) != 0 {
		current_node := queue[0]
		queue = queue[1:]
		for i := 0; i < l; i++ {
			intermediate_word := current_node.word[:i] + "*" + current_node.word[i+1:]
			for j := 0; j < len(dict[intermediate_word]); j++ {
				if dict[intermediate_word][j] == endWord {
					return current_node.state + 1
				}
				if _, ok := set[dict[intermediate_word][j]]; !ok {
					set[dict[intermediate_word][j]] = struct{}{}
					queue = append(queue, node{word: dict[intermediate_word][j], state: current_node.state + 1})
				}
			}
		}
	}
	return 0
}
