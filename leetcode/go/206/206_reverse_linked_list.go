package main

// Problem: https://leetcode.com/problems/reverse-linked-list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	el := head
	for el != nil {
		res = &ListNode{Val: el.Val, Next: res}
		el = el.Next
	}
	return res
}
