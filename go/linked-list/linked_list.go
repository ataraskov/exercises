package linkedlist

import "fmt"

// Define List and Node types here.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
}

func NewList(elements ...interface{}) *List {
	l := List{}
	var prev *Node
	for i, e := range elements {
		n := Node{
			Value: e,
			prev:  prev,
		}
		if prev != nil {
			prev.next = &n
		}
		prev = &n
		if i == 0 {
			l.head = &n
		}
	}
	return &l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

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

func (l *List) Pop() (interface{}, error) {
	last := l.Last()
	if last == nil {
		return nil, fmt.Errorf("empty list")
	}
	if last == l.head {
		l.head = nil
	}
	if last.prev != nil {
		newlast := last.prev
		newlast.next = nil
	}
	return last.Value, nil
}

func (l *List) Reverse() {
	reversed := NewList()
	n := l.head
	for n != nil {
		reversed.Unshift(n.Value)
		n = n.next
	}
	l.head = reversed.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	n := l.head
	for n != nil && n.next != nil {
		n = n.next
	}
	return n
}
