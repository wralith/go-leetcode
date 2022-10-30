package leetcode

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	if n < 2 {
		return n
	}

	l, r, res := 0, 0, 0
	m := make(map[byte]int)

	for r < n {
		m[s[r]] = m[s[r]] + 1

		for len(m) != r-l+1 {
			m[s[l]] = m[s[l]] - 1
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}

	return res
}
