package algo

import (
	"reflect"
)

// Unique modifies a sorted slice so that the unique values are at the beginning
// of the slice. It then returns the length of the unique elements in a slice.
func Unique(slice interface{}, equal func(i, j int) bool) int {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	result := 0
	for first := 1; first < length; first++ {
		if !equal(result, first) {
			result++
			if result != first {
				swap(result, first)
			}
		}
	}
	return result + 1
}
