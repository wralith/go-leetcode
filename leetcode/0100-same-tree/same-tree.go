package leetcode

import "github.com/wralith/go-leetcode/types"

func isSameTree(p *types.TreeNode, q *types.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
