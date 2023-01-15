package leetcode

func longestConsecutive(nums []int) int {
	m, longest := make(map[int]int, len(nums)), 0

	for _, n := range nums {
		if m[n] == 0 {
			lo, hi := findConsecutiveReach(m, n)
			sum := lo + hi + 1

			updateConsecutives(m, n, lo, hi, sum)
			longest = max(longest, sum)
		}
	}

	return longest
}

func findConsecutiveReach(m map[int]int, n int) (lo, hi int) {
	// pre
	if m[n-1] > 0 {
		lo = m[n-1]
	}

	// post
	if m[n+1] > 0 {
		hi = m[n+1]

	}

	return lo, hi
}

// more params...
func updateConsecutives(m map[int]int, n, lo, hi, sum int) {
	m[n] = sum
	m[n-lo] = sum
	m[n+hi] = sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
