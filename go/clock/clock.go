// Pacakge clock implements simple clock type
package clock

import "fmt"

type Clock struct {
	min int
}

func New(h, m int) Clock {
	return Clock{min: (1440 + (h*60+m)%1440) % 1440}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.min+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.min-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}
