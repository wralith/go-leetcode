package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want:     3,
			wantNums: []int{1, 1, 2},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want:     9,
			wantNums: []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}, // Elements after the xx,yy,zz does not matters
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want:     3,
			wantNums: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates2() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums, tt.wantNums) {
				t.Errorf("removeDuplicates2() nums array = %v, want nums %v", tt.args.nums, tt.wantNums)
			}
		})
	}
}
