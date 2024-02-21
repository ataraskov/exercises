package main

// Problem: https://leetcode.com/problems/valid-parentheses

// test cases
// "({[]})" => true
// "(){[]}" => true
// "([){]}" => false
// "(]"     => false
func isValid(s string) bool {
	pairs := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	stack := []rune{}

	for _, r := range s {
		if closing, found := pairs[r]; found {
			stack = append(stack, closing)
			continue
		}
		if r == ')' || r == ']' || r == '}' {
			if len(stack) > 0 && r == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
