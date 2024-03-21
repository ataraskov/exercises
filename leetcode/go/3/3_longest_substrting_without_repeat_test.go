package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5", args{"ab"}, 2},
		{"0", args{""}, 0},
		{"1", args{"bbbbb"}, 1},
		{"2", args{"abcabcbb"}, 3},
		{"3", args{"pwwkew"}, 3},
		{"4", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
