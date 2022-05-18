package main

type Button struct {
	Btn map[byte][]string
	Tmp string
	Ans []string
}

func letterCombinations(digits string) []string {

	if digits == "" {
		return []string{}
	}

	b := new(Button)
	b.Btn = map[byte][]string{'2': {"a", "b", "c"}, '3': {"d", "e", "f"}, '4': {"g", "h", "i"}, '5': {"j", "k", "l"}, '6': {"m", "n", "o"}, '7': {"p", "q", "r", "s"}, '8': {"t", "u", "v"}, '9': {"w", "x", "y", "z"}}

	b.backtrack(digits)

	return b.Ans
}

func (b *Button) backtrack(d string) {

	if len(b.Tmp) == len(d) {
		b.Ans = append(b.Ans, b.Tmp)
		return
	}

	letters := b.Btn[d[len(b.Tmp)]]
	for i := 0; i < len(letters); i++ {
		b.Tmp += letters[i]
		b.backtrack(d)
		b.Tmp = b.Tmp[:len(b.Tmp)-1]
	}

}
