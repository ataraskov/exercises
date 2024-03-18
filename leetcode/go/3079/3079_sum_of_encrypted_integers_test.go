package main

import (
	"testing"
)

func Test_sumOfEncryptedInt(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}}, 6},
		{"2", args{[]int{10, 21, 31}}, 66},
		{"3", args{[]int{109}}, 999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfEncryptedInt(tt.args.nums); got != tt.want {
				t.Errorf("sumOfEncryptedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{523}, 555},
		{"2", args{213}, 333},
		{"3", args{109}, 999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt(tt.args.x); got != tt.want {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
