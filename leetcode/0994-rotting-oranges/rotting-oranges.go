package leetcode

var directions = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func orangesRotting(grid [][]int) int {
	var queue = [][]int{}
	count, fresh := 0, 0

	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				fresh++
			}
			if grid[row][col] == 2 {
				queue = append(queue, []int{row, col})
			}
		}
	}

	for len(queue) > 0 && fresh > 0 {
		n := len(queue) // changes in length shouldn't effect for loop
		for i := 0; i < n; i++ {
			curr := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			for _, direction := range directions {
				row, col := curr[0]+direction[0], curr[1]+direction[1]

				// continue loop if direction is invalid
				if row < 0 || col < 0 || row == len(grid) || col == len(grid[0]) || grid[row][col] != 1 {
					continue
				}

				grid[row][col] = 2
				queue = append(queue, []int{row, col})
				fresh--
			}
		}
		count++
	}

	if fresh == 0 {
		return count
	}

	return -1
}
