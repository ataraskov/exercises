package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var res Ints
	for _, v := range i {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(v int) bool { return !filter(v) })
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var res Lists
	for _, v := range l {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var res Strings
	for _, v := range s {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}
