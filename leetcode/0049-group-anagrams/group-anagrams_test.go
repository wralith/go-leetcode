package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example 1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "Example 2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "Example 3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		got := groupAnagrams(tt.args.strs)

		for _, r := range got {
			sort.Strings(r)
		}

		assert.ElementsMatch(t, got, tt.want)
	}
}
