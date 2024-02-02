// Package binarysearch provides a solution for exercise Binary Search on Exercism's Go Track.
package binarysearch

// SearchInts returns position of the element in given list
func SearchInts(list []int, key int) int {
	start := 0
	end := len(list) - 1
	for start <= end {
		mid := (start + end) / 2
		switch {
		case list[mid] < key:
			start = mid + 1
		case list[mid] > key:
			end = mid - 1
		default:
			return mid
		}
	}
	// not found
	return -1
}
