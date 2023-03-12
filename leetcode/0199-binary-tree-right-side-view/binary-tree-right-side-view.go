package leetcode

import "github.com/wralith/go-leetcode/types"

// BFS - Level Ordered Traversal
// Queue
func rightSideView(root *types.TreeNode) (res []int) {
	q := []*types.TreeNode{root}

	for len(q) != 0 {
		var r *types.TreeNode

		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0] // select leftmost
			q = q[1:]    // remove leftmost
			if node != nil {
				r = node
				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}
		if r != nil {
			res = append(res, r.Val)
		}
	}

	return res
}
