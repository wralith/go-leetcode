package leetcode

import "github.com/wralith/go-leetcode/types"

// Recursive
func lowestCommonAncestor(root, p, q *types.TreeNode) *types.TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}

// // Iterative
// func lowestCommonAncestor(root, p, q *types.TreeNode) *types.TreeNode {
// 	curr := root
//
// 	for curr != nil {
// 		if p.Val > curr.Val && q.Val > curr.Val {
// 			curr = curr.Right
// 		} else if p.Val < curr.Val && q.Val < curr.Val {
// 			curr = curr.Left
// 		} else {
// 			return curr
// 		}
// 	}
//
// 	return nil
// }
