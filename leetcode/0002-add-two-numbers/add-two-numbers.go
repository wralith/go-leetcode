package leetcode

import "github.com/wralith/go-leetcode/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	head := &types.ListNode{Val: 0}
	carry, cur := 0, head
	v1, v2 := 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		l1, v1 = nextNode(l1)
		l2, v2 = nextNode(l2)

		cur.Next = &types.ListNode{Val: (v1 + v2 + carry) % 10}
		cur = cur.Next
		carry = (v1 + v2 + carry) / 10
	}

	return head.Next
}

func nextNode(l *types.ListNode) (*types.ListNode, int) {
	var v int
	if l == nil {
		v = 0
	} else {
		v = l.Val
		l = l.Next
	}

	return l, v
}
