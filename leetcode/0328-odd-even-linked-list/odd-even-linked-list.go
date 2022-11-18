package leetcode

import "github.com/wralith/go-leetcode/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *types.ListNode) *types.ListNode {
	if head == nil {
		return head
	}
	odd, even := head, head.Next
	curOdd, curEven := odd, even

	for curOdd.Next != nil && curEven.Next != nil {
		curOdd.Next = curEven.Next
		curOdd = curOdd.Next

		curEven.Next = curOdd.Next
		curEven = curEven.Next
	}

	curOdd.Next = even

	return head
}
