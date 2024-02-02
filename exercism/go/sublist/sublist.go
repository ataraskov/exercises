// Package sublist implements a function to determine if one list is a sublist of another.
package sublist

// contains checks if l1 contains l2
func contains(l1, l2 []int) bool {
	if len(l1) > len(l2) {
		return false
	}
	if len(l1) == 0 {
		return true
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return contains(l1, l2[1:])
		}
	}
	return true
}

// Sublist checks if one list is a sublist of another.
func Sublist(l1, l2 []int) Relation {
	sublist := contains(l1, l2)
	superlist := contains(l2, l1)

	switch {
	case sublist && len(l1) == len(l2):
		return RelationEqual
	case sublist:
		return RelationSublist
	case superlist:
		return RelationSuperlist
	}

	return RelationUnequal
}
