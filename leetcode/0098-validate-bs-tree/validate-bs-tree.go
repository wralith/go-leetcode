package leetcode

import "github.com/wralith/go-leetcode/types"

func isValidBST(root *types.TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root, min, max *types.TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return dfs(root.Left, min, root) && dfs(root.Right, root, max)
}
