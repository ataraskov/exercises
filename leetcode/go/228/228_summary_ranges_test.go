package main

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]int{}}, []string{}},
		{"2", args{[]int{1}}, []string{"1"}},
		{"3", args{[]int{1, 2, 3}}, []string{"1->3"}},
		{"4", args{[]int{1, 2, 5, 6}}, []string{"1->2", "5->6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
