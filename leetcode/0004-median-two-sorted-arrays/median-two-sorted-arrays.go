package leetcode

import "math"

// It should have O(log(m+n)) time complexity
// Binary Search

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// nums1 always short
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	total := len(nums1) + len(nums2)
	half := (total + 1) >> 1

	var shortLeft, shortRight float64
	var longLeft, longRight float64

	l, r := 0, len(nums1)-1
	for {
		shortSepIdx := (l + r) >> 1
		longSepIdx := half - shortSepIdx - 2

		if shortSepIdx >= 0 {
			shortLeft = float64(nums1[shortSepIdx])
		} else {
			shortLeft = math.Inf(-1)
		}

		if (shortSepIdx + 1) < len(nums1) {
			shortRight = float64(nums1[shortSepIdx+1])
		} else {
			shortRight = math.Inf(1)
		}

		if longSepIdx >= 0 {
			longLeft = float64(nums2[longSepIdx])
		} else {
			longLeft = math.Inf(-1)
		}

		if (longSepIdx + 1) < len(nums2) {
			longRight = float64(nums2[longSepIdx+1])
		} else {
			longRight = math.Inf(1)
		}

		// Result or Binary Search
		if shortLeft <= longRight && longLeft <= shortRight {
			if total%2 == 1 {
				return max(shortLeft, longLeft)
			}
			return (max(shortLeft, longLeft) + min(shortRight, longRight)) / 2
		} else if shortLeft > longRight {
			r = shortSepIdx - 1
		} else {
			l = shortSepIdx + 1
		}
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
