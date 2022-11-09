package leetcode

import "github.com/wralith/go-leetcode/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *types.ListNode) *types.ListNode {
	if head.Next == nil || head == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
