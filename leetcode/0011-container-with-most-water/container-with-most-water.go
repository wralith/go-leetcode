package leetcode

func maxArea(height []int) int {
	f, s, max := 0, len(height)-1, 0
	for s > f {
		area := findArea(height[f], height[s], s-f)
		if area > max {
			max = area
		}

		if height[f] > height[s] {
			s--
		} else {
			f++
		}
	}

	return max
}

func findArea(f, s, distance int) int {
	var shorter int
	if f < s {
		shorter = f
	} else {
		shorter = s
	}

	return distance * shorter
}
