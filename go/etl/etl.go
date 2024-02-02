// Pacakge etl provides a function for transforming a legacy system to a new system.
package etl

import "strings"

// Transform transforms a legacy system to a new system.
func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for k, v := range in {
		for _, s := range v {
			out[strings.ToLower(s)] = k
		}
	}
	return out
}
