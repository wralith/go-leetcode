package leetcode

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	waterCount := 0

	for left < right {
		if maxLeft < maxRight {
			left++
			if maxLeft < height[left] {
				maxLeft = height[left]
			}
			waterCount += areaToFill(maxLeft, maxRight, height[left])
		} else {
			right--
			if maxRight < height[right] {
				maxRight = height[right]
			}
			waterCount += areaToFill(maxLeft, maxRight, height[right])
		}

	}

	return waterCount
}

func areaToFill(left, right, self int) int {
	min := left
	if right < left {
		min = right
	}
	res := min - self

	if res < 0 {
		return 0
	}
	return res
}
