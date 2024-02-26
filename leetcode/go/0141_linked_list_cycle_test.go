package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2, Next: node1}
	node1.Next = node2
	cycled := node1
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has cycle",
			args: args{cycled},
			want: true,
		},
		{
			name: "single node",
			args: args{&ListNode{Val: 1}},
			want: false,
		},
		{
			name: "no cycle",
			args: args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
