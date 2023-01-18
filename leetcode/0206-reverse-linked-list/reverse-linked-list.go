package leetcode

import "github.com/wralith/go-leetcode/types"

// Iterative: Time O(n), Memory O(1)
func reverseList(head *types.ListNode) *types.ListNode {
	var prev *types.ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// Recursive: Time O(n), Memory O(n)
func reverseListRecursive(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// Recursive: Alternative with helper func
func reverseListRecursiveHelper(head *types.ListNode) *types.ListNode {
	return recurse(head, nil)
}

func recurse(curr, prev *types.ListNode) *types.ListNode {
	if curr == nil {
		return prev
	}

	next := curr.Next
	curr.Next = prev
	return recurse(next, curr)
}
