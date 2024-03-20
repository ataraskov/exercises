package main

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"4", args{"a", "A"}, false},
		{"0", args{"", ""}, true},
		{"1", args{"abc", "abc"}, true},
		{"2", args{"ab#c", "ad#c"}, true},
		{"3", args{"ab##", "c#d#"}, true},
		{"5", args{"gt##", "g#"}, true},
		{"6", args{"abc#d##", "a"}, true},
		{"7", args{"ab##", "c#d#"}, true},
		{"8", args{"bbbextm", "bbb#extm"}, false},
		{"9", args{"a##c", "#a#c"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
