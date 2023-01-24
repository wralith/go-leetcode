package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMedianSortedArrays(t *testing.T) {
	n1 := []int{1, 3}
	n2 := []int{2}
	got := findMedianSortedArrays(n1, n2)
	want := 2.0
	require.Equal(t, got, want)

	n1 = []int{1, 2}
	n2 = []int{3, 4}
	got = findMedianSortedArrays(n1, n2)
	want = 2.5
	require.Equal(t, got, want)

	n1 = []int{}
	n2 = []int{2, 3}
	got = findMedianSortedArrays(n1, n2)
	want = 2.5
	require.Equal(t, got, want)
}
