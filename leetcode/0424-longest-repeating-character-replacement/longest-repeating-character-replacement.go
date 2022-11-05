package leetcode

func characterReplacement(s string, k int) int {
	l, r, res, freq := 0, 0, 0, 0
	hm := make(map[byte]int)

	for r < len(s) {
		hm[s[r]]++
		freq = max(freq, hm[s[r]])

		if (r-l+1)-freq > k {
			hm[s[l]]--
			l++
		}

		res = max(res, r-l+1)
		r++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
