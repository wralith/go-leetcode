package leetcode

import "github.com/wralith/go-leetcode/types"

func mergeTwoLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	list := &types.ListNode{}
	head := list

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}

	return list.Next
}
