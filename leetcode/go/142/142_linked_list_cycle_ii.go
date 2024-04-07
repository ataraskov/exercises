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
	/*
		seen := make(map[*ListNode]bool)
		for head != nil {
			if seen[head] {
				return head
			}
			seen[head] = true
			head = head.Next
		}
	*/

	// two pointers
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(1)
	slow := head
	fast := head
	isCycle := false
	// search for a loop
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	// find looped el
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
