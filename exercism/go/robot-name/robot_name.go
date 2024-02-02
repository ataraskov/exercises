// package robotname provides a colution for Robot Name on Exercism's Go Track exercise
package robotname

import (
	"fmt"
)

var (
	initialized = false
	names       = map[string]bool{}
	letters     = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type Robot struct {
	name string
}

// populateNames generates all possible names in advance
func populateNames() {
	for l1 := 0; l1 < len(letters); l1++ {
		for l2 := 0; l2 < len(letters); l2++ {
			for i := 0; i <= 999; i++ {
				s := fmt.Sprintf(
					"%s%s%03d",
					string(letters[l1]),
					string(letters[l2]),
					i,
				)
				names[s] = true
			}
		}
	}
	initialized = true
}

// Name returns robot's name
func (r *Robot) Name() (string, error) {
	if !initialized {
		populateNames()
	}
	if r.name == "" {
		if len(names) == 0 {
			return "", fmt.Errorf("namespace exhausted")
		}
		for k := range names {
			r.name = k
		}
		delete(names, r.name)
	}
	return r.name, nil
}

// Reset, gives a new name to the robot
func (r *Robot) Reset() error {
	r.name = ""
	_, err := r.Name()
	return err
}
