package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1}, 1}, 0},
		{"2", args{[]int{1, 2, 3}, 2}, 1},
		{"3", args{[]int{1, 2, 3}, 4}, -1},
		{"4", args{[]int{1, 2, 3}, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
