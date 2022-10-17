package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) (res [][]string) {
	m := make(map[string][]string, len(strs))

	for _, s := range strs {
		arr := []rune(s)
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		m[string(arr)] = append(m[string(arr)], s)
	}

	for _, s := range m {
		res = append(res, s)
	}

	return res
}
