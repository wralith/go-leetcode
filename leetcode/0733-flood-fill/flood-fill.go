package leetcode

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	startingPoint := image[sr][sc]
	if startingPoint == color {
		return image
	}

	dfs(image, sr, sc, color, startingPoint)

	return image
}

func dfs(image [][]int, sr int, sc int, color int, prevColor int) {

	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != prevColor {
		return
	}

	image[sr][sc] = color
	dfs(image, sr+1, sc, color, prevColor)
	dfs(image, sr-1, sc, color, prevColor)
	dfs(image, sr, sc+1, color, prevColor)
	dfs(image, sr, sc-1, color, prevColor)

}
