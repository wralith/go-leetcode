package types

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListFromInts(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	l2 := l
	for _, v := range nums {
		l2.Next = &ListNode{Val: v}
		l2 = l2.Next
	}

	return l.Next
}
