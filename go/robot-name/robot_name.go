// package robotname provides a colution for Robot Name on Exercism's Go Track exercise
package robotname

import (
	"fmt"
	"math/rand"
)

var (
	initialized = false
	names       = []string{}
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
				names = append(names, s)
			}
		}
	}
	initialized = true
}

// Name returns robot's name
// TODO: very inefficient, need to use a map with delete op instead
func (r *Robot) Name() (string, error) {
	if !initialized {
		populateNames()
	}
	if r.name == "" {
		i := rand.Intn(len(names))
		r.name = names[i]
		names = append(names[:i], names[i+1:]...)
	}
	return r.name, nil
}

// Reset, gives a new name to the robot
func (r *Robot) Reset() error {
	r.name = ""
	_, err := r.Name()
	return err
}
