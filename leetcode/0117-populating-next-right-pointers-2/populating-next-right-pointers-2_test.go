package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_connect(t *testing.T) {
	x4 := types.TreeNodeWithNext{Val: 4}
	x5 := types.TreeNodeWithNext{Val: 5}
	x7 := types.TreeNodeWithNext{Val: 7}
	x3 := types.TreeNodeWithNext{Val: 3, Right: &x7}
	x2 := types.TreeNodeWithNext{Val: 2, Left: &x4, Right: &x5}
	x1 := types.TreeNodeWithNext{Val: 1, Left: &x2, Right: &x3}

	connect(&x1)

	require.Equal(t, x3.Val, x2.Next.Val)
	require.Equal(t, x5.Val, x4.Next.Val)
	require.Equal(t, x7.Val, x5.Next.Val)
	require.Nil(t, x1.Next)
	require.Nil(t, x3.Next)
	require.Nil(t, x7.Next)
}
