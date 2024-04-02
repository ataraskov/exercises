package main

import (
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"1", args{nil}, nil},
		{"2", args{&Node{Val: 1, Child: &Node{Val: 2, Child: &Node{Val: 3}}}},
			&Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}},
		{"3", args{&Node{Val: 1, Next: &Node{Val: 2}, Child: &Node{Val: 3}}},
			&Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 2}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := flatten(tt.args.root)
			for {
				switch {
				case got == nil && want != nil:
					t.Errorf("flatten() = %v, want %v (short)", got, tt.want)
				case got != nil && want == nil:
					t.Errorf("flatten() = %v, want %v (long)", got, tt.want)
				case got == nil && want == nil:
					break
				case got.Val != want.Val:
					t.Errorf("flatten() = %v, want %v (value)", got, tt.want)
				}
				if got == nil || want == nil {
					break
				}
				want = want.Next
				got = got.Next

			}
		})
	}
}
