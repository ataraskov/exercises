// Package flatten provides solution for Flatten Array on Exercism's Go Track
package flatten

import (
	"reflect"
)

// Flatten takes a nested list ignores all nil values and returns flattened slice
func Flatten(nested interface{}) []interface{} {
	res := make([]interface{}, 0)
	n := reflect.ValueOf(nested)
	if n.Kind() == reflect.Slice {
		for i := 0; i < n.Len(); i++ {
			switch {
			case n.Index(i).IsNil():
				continue
			case reflect.ValueOf(n.Index(i).Interface()).Kind() == reflect.Slice:
				res = append(res, Flatten(n.Index(i).Interface())...)
			default:
				res = append(res, n.Index(i).Interface())
			}
		}
	}
	return res
}
