package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}},
		{"2", args{&ListNode{1, nil}}, &ListNode{1, nil}},
		{"3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
