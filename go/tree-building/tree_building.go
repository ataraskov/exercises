// Package tree contains a solution for the Tree Building on Exercism's Go Track
package tree

import (
	"errors"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

var (
	ErrCycle         = errors.New("records cycle detected")
	ErrRootMissing   = errors.New("no root node found")
	ErrRootParentId  = errors.New("no root node found")
	ErrParentId      = errors.New("incorrect parrent id")
	ErrDuplicate     = errors.New("duplicate record")
	ErrDuplicateRoot = errors.New("duplicate root record")
	ErrMissing       = errors.New("missing a record")
)

// Build returns a tree structure of Nodes from given Record slice
// Be adviced: this is a really bad implementaion
func Build(records []Record) (*Node, error) {
	if len(records) < 1 {
		return nil, nil
	}

	nodes := make([]Node, len(records))

	// requirement: The ID number is always between 0 (inclusive)
	//              and the length of the record list (exclusive)
	for i := 0; i < len(records); i++ {
		found := false

		for _, r := range records {
			if r.ID > 0 && r.ID <= r.Parent {
				return nil, ErrParentId
			}
			if i == r.ID {
				if r.ID == 0 && r.Parent != 0 {
					return nil, ErrRootParentId
				}

				// node
				nodes[r.ID] = Node{ID: r.ID}

				// children
				parent := &nodes[r.Parent]
				child := &nodes[r.ID]
				if parent != child {
					parent.Children = append(parent.Children, child)
				}

				found = true
				break
			}

		}

		if i == 0 && !found {
			return nil, ErrRootMissing
		}
		if !found {
			return nil, ErrMissing
		}
	}
	return &nodes[0], nil
}
