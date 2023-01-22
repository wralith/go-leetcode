package leetcode

import (
	"github.com/wralith/go-leetcode/types"
)

func isBalanced(root *types.TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *types.TreeNode) (balanced bool, height int) {
	if root == nil {
		return true, 0
	}
	left, leftHeight := dfs(root.Left)
	right, rightHeight := dfs(root.Right)

	balanced = left && right && abs(leftHeight-rightHeight) <= 1
	height = 1 + max(leftHeight, rightHeight)

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// or math.Abs
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
