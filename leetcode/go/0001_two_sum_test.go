package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{[]int{0, 0, 1, 1}, 2},
			[]int{2, 3},
		},
		{
			args{[]int{1, 1}, 2},
			[]int{0, 1},
		},
		{
			args{[]int{1, 0, 0, 1}, 2},
			[]int{0, 3},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
