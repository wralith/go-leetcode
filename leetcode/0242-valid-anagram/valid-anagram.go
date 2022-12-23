package leetcode

import (
	"bytes"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := []byte(s)
	s2 := []byte(t)

	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	return bytes.Equal(s1, s2)
}
