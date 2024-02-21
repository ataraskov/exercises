package main

// Problem: https://leetcode.com/problems/same-tree

// test cases
// nil		  nil        => true
// [1,2,3]    [1,2,3]    => true
// [1,2,null] [1,null,2] => false
// [1,2,1]    [1,1,2]    => false

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
