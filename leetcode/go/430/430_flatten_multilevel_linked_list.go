package main

// Problem: https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	el := root
	stack := []*Node{}
	for el != nil {
		var next *Node
		if el.Child != nil {
			if el.Next != nil {
				stack = append(stack, el.Next)
			}
			el.Next = el.Child
			el.Child.Prev = el
			next = el.Child
			el.Child = nil
		} else {
			next = el.Next
		}
		if next == nil && len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			el.Next = tmp
			tmp.Prev = el
			next = tmp
		}
		el = next
	}
	return root
}
