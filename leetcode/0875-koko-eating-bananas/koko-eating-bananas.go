package leetcode

import "math"

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, maxIn(piles)
	min := hi

	for lo < hi {
		mid := lo + (hi-lo)>>1
		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(mid)))
		}

		if hours <= h {
			min = minimum(min, mid)
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return min
}

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxIn(a []int) (res int) {
	for _, x := range a {
		if x > res {
			res = x
		}
	}
	return
}
