// Package linkedlist provides a solution for exercise Simple Linked List on Exercism's Go Track.
package linkedlist

import "fmt"

// Node represents a node of linked list.
type Element struct {
	Value int
	next  *Element
}

// List represents a linked list containing Elements.
type List struct {
	head *Element
}

// New creates a linked list from given slice
func New(elements []int) *List {
	var head *Element
	var prev *Element
	for _, e := range elements {
		elem := &Element{Value: e}
		if head == nil {
			head = elem
		} else {
			prev.next = elem
		}
		prev = elem
	}
	return &List{head: head}
}

// Size returns count of elements
func (l *List) Size() int {
	size := 0
	elem := l.head

	for elem != nil {
		size++
		if elem.next == nil {
			break
		}
		elem = elem.next
	}
	return size
}

// Push adds an element to the tail
func (l *List) Push(element int) {
	elem := l.head
	for elem != nil {
		if elem.next == nil {
			elem.next = &Element{Value: element}
			return
		}
		elem = elem.next
	}
	l.head = &Element{Value: element}
}

// Pop removes the tail element
func (l *List) Pop() (int, error) {
	elem := l.head
	if l == nil || l.head == nil {
		return 0, fmt.Errorf("empty list")
	}

	var prev *Element
	var res int
	for elem != nil {
		if elem.next == nil {
			res = elem.Value
			if prev != nil {
				prev.next = nil
			} else {
				l.head = nil
			}
		}
		prev = elem
		elem = elem.next
	}
	return res, nil
}

// Array returns an array of elements
func (l *List) Array() []int {
	arr := []int{}
	if l != nil {
		elem := l.head

		for elem != nil {
			arr = append(arr, elem.Value)
			if elem.next == nil {
				break
			}
			elem = elem.next
		}
	}
	return arr
}

// Reverse returns a reversed list
func (l *List) Reverse() *List {
	arr := []int{}
	if l == nil || l.head == nil {
		return &List{}
	}
	elem := l.head

	for elem != nil {
		arr = append(arr, elem.Value)
		if elem.next == nil {
			break
		}
		elem = elem.next
	}

	list := &List{}
	for i := len(arr) - 1; i >= 0; i-- {
		list.Push(arr[i])
	}
	return list
}
