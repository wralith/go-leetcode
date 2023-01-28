package leetcode

import "github.com/wralith/go-leetcode/types"

func maxDepth(root *types.TreeNode) (result int) {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
