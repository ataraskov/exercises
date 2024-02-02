package stringset

import (
	"bytes"
	"fmt"
)

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, i := range l {
		s[i] = true
	}
	return s
}

// String formats set as string.
// For example, a set with 2 elements, "a" and "b", will be formatted as {"a", "b"}.
func (s Set) String() string {
	buf := bytes.Buffer{}
	for k := range s {
		if buf.Len() > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%q", k))
	}
	return fmt.Sprintf("%s%s%s", "{", buf.String(), "}")
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	s := New()

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for v := range s1 {
		if s2.Has(v) {
			s.Add(v)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for v := range s1 {
		if !s2.Has(v) {
			s.Add(v)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for v := range s1 {
		s.Add(v)
	}
	for v := range s2 {
		s.Add(v)
	}
	return s
}
