package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test case 1", args{[]int{}, 1}, 0},
		{"test case 1", args{[]int{1, 1}, 1}, 0},
		{"test case 2", args{[]int{1, 1, 2}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
