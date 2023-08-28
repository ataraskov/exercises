// Package tree contains a solution for the Tree Building on Exercism's Go Track
package tree

import (
	"errors"
	"sort"
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
	ErrRoot     = errors.New("no root node found")
	ErrParentId = errors.New("incorrect parrent id")
	ErrMissing  = errors.New("missing a record")
)

// Build returns a tree structure of Nodes from given Record slice
func Build(records []Record) (*Node, error) {
	if len(records) < 1 {
		return nil, nil
	}

	nodes := make([]Node, len(records))

	// sort input
	sorted := make([]Record, len(records))
	copy(sorted, records)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ID < sorted[j].ID
	})

	// requirement: The ID number is always between 0 (inclusive)
	//              and the length of the record list (exclusive)
	for i, r := range sorted {
		if r.ID > 0 && r.ID <= r.Parent {
			return nil, ErrParentId
		}
		if r.ID == 0 && r.Parent != 0 {
			return nil, ErrRoot
		}
		if i != r.ID {
			return nil, ErrMissing
		}

		// node
		nodes[r.ID] = Node{ID: r.ID}
		if r.ID != 0 {
			// children
			parent := &nodes[r.Parent]
			child := &nodes[r.ID]
			parent.Children = append(parent.Children, child)
		}

	}
	return &nodes[0], nil
}
