package main

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 2, 0}}, 3},
		{"case2", args{[]int{3, 4, -1, 1}}, 2},
		{"case3", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"case4", args{[]int{1}}, 2},
		{"case5", args{[]int{-1, 4, 2, 1, 9, 10}}, 3},
		{"case6", args{[]int{1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
