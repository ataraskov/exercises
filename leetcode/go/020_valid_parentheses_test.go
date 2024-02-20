package main

import (
	"fmt"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args{"({[]})"},
			true,
		},
		{
			args{"(){[]}"},
			true,
		},
		{
			args{"([){]}"},
			false,
		},
		{
			args{"(]"},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
