package leetcode

func numIslands(grid [][]byte) (result int) {
	m, n := len(grid), len(grid[0]) // Vertical and horizontal lengths

	visited := make([][]bool, m) // create a visited slice [coordinate: bool]
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ { // X Axis
		for j := 0; j < n; j++ { // Y Axis
			if grid[i][j] == '1' && !visited[i][j] { // If it is land and not visited yet
				result++                  // We discovered a new island!
				dfs(i, j, grid, &visited) // Search the edges of the island
			}
		}
	}

	return
}

func dfs(row, col int, grid [][]byte, visited *[][]bool) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || (*visited)[row][col] || grid[row][col] == '0' {
		return // return if in border, if visited or if sea
	}

	(*visited)[row][col] = true // mark as visited

	dfs(row+1, col, grid, visited) // search down
	dfs(row-1, col, grid, visited) // search up
	dfs(row, col+1, grid, visited) // search right
	dfs(row, col-1, grid, visited) // search left
}
