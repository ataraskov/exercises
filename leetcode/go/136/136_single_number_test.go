package main

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[2,2,1]", args{[]int{2, 2, 1}}, 1},
		{"[1,2,2]", args{[]int{1, 2, 2}}, 1},
		{"[2,1,2]", args{[]int{2, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
