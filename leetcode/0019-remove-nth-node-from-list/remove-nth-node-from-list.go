package leetcode

import "github.com/wralith/go-leetcode/types"

func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return slow.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}

// Another approach, find length in first traversal, remove len-n'th element
func removeNthFromEnd2(head *types.ListNode, n int) *types.ListNode {
	ptr := head
	length := 0

	for ptr != nil {
		ptr = ptr.Next
		length++
	}

	if length == 1 {
		return nil
	} else if length == n {
		return head.Next
	}

	ptr = head
	for i := 1; i < length-n; i++ {
		ptr = ptr.Next
	}

	ptr.Next = ptr.Next.Next
	return head
}
