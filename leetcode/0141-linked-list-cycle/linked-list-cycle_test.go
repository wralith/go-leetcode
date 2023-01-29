package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_hasCycle(t *testing.T) {
	head := types.MakeListFromInts(3, 2, 0, -4)
	head.Next.Next.Next = head
	got := hasCycle(head)
	require.True(t, got)

	head = types.MakeListFromInts(1, 2)
	head.Next.Next = head
	got = hasCycle(head)
	require.True(t, got)

	head = types.MakeListFromInts(1)
	got = hasCycle(head)
	require.False(t, got)
}
