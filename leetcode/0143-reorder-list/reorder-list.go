package leetcode

import "github.com/wralith/go-leetcode/types"

func reorderList(head *types.ListNode) {
	// Slow will be half
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half
	secondHalf := reverse(slow.Next)
	slow.Next = nil // Unlink

	firstHalf := head

	for firstHalf != nil && secondHalf != nil {
		firstHalf, firstHalf.Next, secondHalf, secondHalf.Next = firstHalf.Next, secondHalf, secondHalf.Next, firstHalf.Next
	}
}

// 206 - Reverse Linked List
func reverse(head *types.ListNode) *types.ListNode {
	var prev *types.ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
