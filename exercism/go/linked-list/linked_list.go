// Pacakge linkedlist provides a solution for exercise Linked List on Exercism's Go Track.
package linkedlist

import "fmt"

// Node represents a node of linked list.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

// List represents a linked list containing Nodes.
type List struct {
	head *Node
}

// NewList creates a new linked list from givin elements.
func NewList(elements ...interface{}) *List {
	l := List{}
	for _, e := range elements {
		l.Push(e)
	}
	return &l
}

// Next returns a pointer to the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns a pointer to the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift inserts a value at the front of the list.
func (l *List) Unshift(v interface{}) {
	n := &Node{
		Value: v,
	}
	if l.head != nil {
		l.head.prev = n
	}
	n.next = l.head
	l.head = n
}

// Push inserts a value at the back of the list.
func (l *List) Push(v interface{}) {
	n := &Node{
		Value: v,
	}
	last := l.Last()
	if last != nil {
		last.next = n
	} else {
		l.head = n
	}
	n.prev = last
}

// Shift removes the head node from the list.
func (l *List) Shift() (interface{}, error) {
	if l.head != nil {
		old := l.head
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		return old.Value, nil
	}
	return nil, fmt.Errorf("empty list")
}

// Pop removes last node from the list.
func (l *List) Pop() (interface{}, error) {
	last := l.Last()

	switch last {
	case nil:
		return nil, fmt.Errorf("empty list")
	case l.head:
		l.head = nil
	default:
		last.prev.next = nil
	}

	return last.Value, nil
}

// Reverse updates node lins to reverse the list.
func (l *List) Reverse() {
	reversed := NewList()
	n := l.head
	for n != nil {
		reversed.Unshift(n.Value)
		n = n.next
	}
	l.head = reversed.head
}

// First returns a pointer to the head node of the list.
func (l *List) First() *Node {
	return l.head
}

// First returns a pointer to the last node of the list.
func (l *List) Last() *Node {
	n := l.head
	for n != nil && n.next != nil {
		n = n.next
	}
	return n
}
