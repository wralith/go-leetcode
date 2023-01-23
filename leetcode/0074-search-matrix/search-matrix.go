package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix)
	colLen := len(matrix[0])

	// Go to the middle column, check if end element is greater or lesser than target, BS
	start, end := 0, rowLen-1
	for start <= end {
		mid := (start + end) / 2
		if target == matrix[mid][0] || target == matrix[mid][colLen-1] {
			return true
		}
		// BS in the row
		if target > matrix[mid][0] && target < matrix[mid][colLen-1] {
			return bs(matrix[mid], target)
		} else if target > matrix[mid][colLen-1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}

func bs(slice []int, target int) bool {
	start, end := 0, len(slice)-1
	for start <= end {
		mid := (start + end) / 2
		if slice[mid] == target {
			return true
		} else if slice[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
