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

func preorderTraversal(root *types.TreeNode) []int {
	var res []int
	preorder(root, &res) // We pass reference to res, so preorder function will mutate the res slice
	return res
}

func preorder(root *types.TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preorder(root.Left, res)
		preorder(root.Right, res)
	}
}
