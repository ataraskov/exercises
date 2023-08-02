package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	children []string
	plants   [][]string
}

var Plants = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !strings.HasPrefix(diagram, "\n") {
		return nil, errors.New("wrong diagram format")
	}

	childrenCopy := make([]string, len(children))
	copy(childrenCopy, children)
	duplicate := false
	sort.Slice(childrenCopy, func(i, j int) bool {
		switch strings.Compare(childrenCopy[i], childrenCopy[j]) {
		case 0:
			duplicate = true
		case -1:
			return true
		default:
		}
		return false
	})
	if duplicate {
		return nil, errors.New("duplicate name")
	}

	rows := strings.Split(strings.TrimPrefix(diagram, "\n"), "\n")
	plants := [][]string{
		{},
		{},
	}
	for r_index, r := range rows {
		if r_index > 0 && len(rows[r_index-1]) != len(rows[r_index]) {
			return nil, errors.New("rows mismatch")
		}

		if len(r) != len(childrenCopy)*2 {
			return nil, errors.New("caps to children mismatch")
		}

		for _, c := range r {
			if plant, ok := Plants[string(c)]; ok {
				plants[r_index] = append(plants[r_index], plant)
			} else {
				return nil, errors.New("invalid plant code")
			}
		}
	}
	return &Garden{children: childrenCopy, plants: plants}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	index := -1
	for i, c := range g.children {
		if c == child {
			index = i * 2
		}
	}

	if index < 0 {
		return []string{}, false
	}

	res := []string{}
	for _, r := range g.plants {
		res = append(res, r[index])
		res = append(res, r[index+1])
	}
	return res, true
}
