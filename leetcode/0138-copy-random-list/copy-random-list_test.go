package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_copyRandomList(t *testing.T) {
	head := MakeListFromSlice([][2]int{
		{7, types.NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0},
	})
	got := copyRandomList(head)
	require.Equal(t, head, got)
	require.NotSame(t, head, got)

	// Example 2
	head = MakeListFromSlice([][2]int{
		{1, 1}, {2, 1},
	})
	got = copyRandomList(head)
	require.Equal(t, head, got)
	require.NotSame(t, head, got)

	// Example 3
	head = MakeListFromSlice([][2]int{
		{1, 1}, {2, 1},
	})
	got = copyRandomList(head)
	require.Equal(t, head, got) // Example 2
	head = MakeListFromSlice([][2]int{
		{3, types.NULL}, {3, 0}, {3, types.NULL},
	})
	got = copyRandomList(head)

	require.Equal(t, head, got)
	require.NotSame(t, head, got)
	require.NotSame(t, head, got)
}

func MakeListFromSlice(s [][2]int) *Node {
	if len(s) == 0 {
		return nil
	}
	nodesArr := []*Node{}
	head := &Node{}
	curr := head
	for _, v := range s {
		curr.Next = &Node{Val: v[0]}
		nodesArr = append(nodesArr, curr)
		curr = curr.Next
	}

	curr = head
	for _, v := range s {
		if v[1] == types.NULL {
			continue
		}
		curr.Random = nodesArr[v[1]]
		curr = curr.Next
	}

	return head
}
