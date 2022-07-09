package main

func generateParenthesis(n int) []string {
	res := genParent([]string{}, "", n, n)
	return res
}

func genParent(res []string, s string, open, close int) []string {
	if close == 0 {
		res = append(res, s)
	}

	if open > 0 {
		res = genParent(res, s+"(", open-1, close)
	}

	if open < close {
		res = genParent(res, s+")", open, close-1)
	}

	return res
}
