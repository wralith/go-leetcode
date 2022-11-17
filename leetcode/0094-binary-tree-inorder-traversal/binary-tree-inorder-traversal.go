package leetcode

import "github.com/wralith/go-leetcode/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *types.TreeNode) []int {
	var res []int
	dfs(root, &res)
	return res
}

func dfs(root *types.TreeNode, res *[]int) {
	if root != nil {
		dfs(root.Left, res)
		*res = append(*res, root.Val)
		dfs(root.Right, res)
	}
}
