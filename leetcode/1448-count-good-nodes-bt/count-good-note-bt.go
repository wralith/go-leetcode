package leetcode

import "github.com/wralith/go-leetcode/types"

func goodNodes(root *types.TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *types.TreeNode, parent int) int {
	if root == nil {
		return 0
	}

	max, res := maxAndBit(root.Val, parent)

	res += dfs(root.Left, max)
	res += dfs(root.Right, max)

	return res
}

// Weird yeah, returns max number itself and 1 if first arg is greater of equal, private 🤫
func maxAndBit(a, b int) (int, int) {
	if a >= b {
		return a, 1
	}
	return b, 0
}
