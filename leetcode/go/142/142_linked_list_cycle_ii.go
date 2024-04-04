package main

// Problem: https://leetcode.com/problems/linked-list-cycle-ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(n)
	seen := make(map[*ListNode]bool)
	for head != nil {
		if seen[head] {
			return head
		}
		seen[head] = true
		head = head.Next
	}
	return head
}
