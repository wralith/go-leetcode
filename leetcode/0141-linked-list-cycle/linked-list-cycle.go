package leetcode

import "github.com/wralith/go-leetcode/types"

func hasCycle(head *types.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	f, s := head, head
	for s != nil && s.Next != nil {
		f = f.Next
		s = s.Next.Next
		if f == s {
			return true
		}
	}
	return false
}
