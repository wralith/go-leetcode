package leetcode

import "github.com/wralith/go-leetcode/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *types.ListNode) *types.ListNode {
	a, b := headA, headB
	if a == nil || b == nil {
		return nil
	}
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
