package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_evalRPN(t *testing.T) {
	want := 9
	got := evalRPNSlice([]string{"2", "1", "+", "3", "*"})
	gotList := evalRPN([]string{"2", "1", "+", "3", "*"})
	require.Equal(t, want, got)
	require.Equal(t, want, gotList)

	want = 6
	got = evalRPNSlice([]string{"4", "13", "5", "/", "+"})
	gotList = evalRPN([]string{"4", "13", "5", "/", "+"})
	require.Equal(t, want, got)
	require.Equal(t, want, gotList)

	want = 22
	got = evalRPNSlice([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	gotList = evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	require.Equal(t, want, got)
	require.Equal(t, want, gotList)
}

func Benchmark_evalRPN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	}
}

func Benchmark_evalRPNSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evalRPNSlice([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	}
}
