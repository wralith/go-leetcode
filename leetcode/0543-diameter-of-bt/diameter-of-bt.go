package leetcode

import "github.com/wralith/go-leetcode/types"

func diameterOfBinaryTree(root *types.TreeNode) (result int) {
	dfs(root, &result)
	return
}

func dfs(t *types.TreeNode, maxLen *int) int {
	if t == nil {
		return 0
	}

	l := dfs(t.Left, maxLen)
	r := dfs(t.Right, maxLen)
	*maxLen = max(*maxLen, l+r)

	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
