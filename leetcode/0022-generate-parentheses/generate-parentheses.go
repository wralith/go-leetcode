package leetcode

func generateParenthesis(n int) []string {
	res := []string{}
	backTrack(n, n, "", &res)
	return res
}

// open, close => remaining open or close parentheses to add
// close should be less than open and equal to open in result
func backTrack(open, close int, s string, res *[]string) {
	if open == 0 && close == 0 {
		*res = append(*res, s)
		return
	}
	if open > 0 {
		backTrack(open-1, close, s+"(", res)
	}
	if close > 0 && close > open {
		backTrack(open, close-1, s+")", res)
	}
}
