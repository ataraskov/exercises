// Package listops implements operations on lists.
package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Foldl redues a list from the left by applying a function(accumulator, item) to each element of the list
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	res := initial
	for _, el := range s {
		res = fn(res, el)
	}
	return res
}

// Foldr redues a list from the right by applying a function(item, accumulator) to each element of the list
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	res := initial
	for _, el := range s.Reverse() {
		res = fn(el, res)
	}
	return res

}

// Filter returns a list of elements that satisfy given predicate
func (s IntList) Filter(fn func(int) bool) IntList {
	res := IntList{}
	for _, el := range s {
		if fn(el) {
			res = res.Append(IntList{el})
		}
	}
	return res
}

// Length returns the length of a list
func (s IntList) Length() int {
	return s.Foldl(func(x, y int) int { return x + 1 }, 0)
}

// Map returns a list after applying given function to each element of the list
func (s IntList) Map(fn func(int) int) IntList {
	res := IntList{}
	for _, i := range s {
		res = append(res, fn(i))
	}
	return res
}

// Reverse reverses a list
func (s IntList) Reverse() IntList {
	res := make(IntList, s.Length())
	for i, el := range s {
		res[s.Length()-i-1] = el
	}
	return res
}

// Append adds a list to another
func (s IntList) Append(lst IntList) IntList {
	res := make(IntList, s.Length()+lst.Length())
	copy(res, s)
	for i, el := range lst {
		res[s.Length()+i] = el
	}
	return res
}

// Concat merges lists together
func (s IntList) Concat(lists []IntList) IntList {
	res := make(IntList, s.Length())
	copy(res, s)
	for _, l := range lists {
		res = res.Append(l)
	}
	return res
}
