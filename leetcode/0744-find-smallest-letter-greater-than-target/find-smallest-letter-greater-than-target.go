package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	l := len(letters)
	left, right := 0, l-1

	if letters[l-1] <= target || target < letters[0] {
		return letters[0]
	}

	for left+1 < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	return letters[right]
}
