package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	l := ""
	if len(strs) == 0 {
		return l
	}

	l = strs[0]

	for _, s := range strs {
		for !strings.HasPrefix(s, l) {
			l = l[0:(len(l) - 1)]
		}
	}

	return l
}
