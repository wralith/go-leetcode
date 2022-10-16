package leetcode

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	var digitLogs []string
	var letterLogs []string

	distributeLogsByType(&logs, &digitLogs, &letterLogs)

	sort.SliceStable(letterLogs, func(i, j int) bool {
		s1 := strings.SplitN(letterLogs[i], " ", 2) // {head, tail}
		s2 := strings.SplitN(letterLogs[j], " ", 2)
		// If tails are equal compare heads
		if s1[1] == s2[1] {
			return s1[0] < s2[0]
		}
		return s1[1] < s2[1]
	})

	res := append(letterLogs, digitLogs...)
	return res
}

// Decided to extract this logic out of main function
func distributeLogsByType(logs *[]string, digitLogs *[]string, letterLogs *[]string) {
	for _, log := range *logs {
		l := strings.Split(log, " ")
		if isDigit(l[1]) {
			*digitLogs = append(*digitLogs, log)
		} else {
			*letterLogs = append(*letterLogs, log)
		}
	}
}

// Helper from the optimal solution in leetcode, seems better than unicode.IsDigit
// additionally strconv.atoi did not work as i thought at big numbers
func isDigit(s string) bool { return s[0] >= '0' && s[0] <= '9' }
