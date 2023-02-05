package leetcode

import "github.com/wralith/go-leetcode/types"

// Recursive
func kthSmallest(root *types.TreeNode, k int) int {
	_, res := traverseInOrder(root, k, 0, 0)
	return res
}

func traverseInOrder(node *types.TreeNode, k, c, res int) (int, int) {
	if node != nil {
		// Go to left until there is none
		c, res = traverseInOrder(node.Left, k, c, res)

		c++
		// K'th smallest element, this is our result!
		if c == k {
			res = node.Val
			return c, res
		}

		c, res = traverseInOrder(node.Right, k, c, res)
	}
	return c, res
}

// Iterative
func kthSmallest2(root *types.TreeNode, k int) int {
	stack := []*types.TreeNode{}

	for {
		// Traverse left child until there is none
		for root != nil {
			// Push to stack
			stack = append(stack, root)
			root = root.Left
		}

		// Pop from stack
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}
