package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	cycle := &ListNode{Val: 8}
	cycle.Next = cycle
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{nil}, nil},
		{"2", args{&ListNode{Val: 1}}, nil},
		{"2", args{cycle}, cycle},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
