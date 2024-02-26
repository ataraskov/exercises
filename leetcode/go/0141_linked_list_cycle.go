package main

// Problem: https://leetcode.com/problems/linked-list-cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if seen[node] {
			return true
		}
		seen[node] = true
		node = node.Next
	}
	return false
}
