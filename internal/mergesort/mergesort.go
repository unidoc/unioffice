// Package mergesort provides an implementation for merge sort sorting algorithm
package mergesort

// MergeSort receives a slice and returns a new sorted slice
func MergeSort(array []float64) []float64 {
	if len(array) <= 1 {
		result := make([]float64, len(array))
		copy(result, array)
		return result
	}

	mid := len(array) / 2
	leftArray := MergeSort(array[:mid])
	rightArray := MergeSort(array[mid:])

	result := make([]float64, len(array))

	arrayIndex := 0
	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(leftArray) && rightIndex < len(rightArray) {
		if leftArray[leftIndex] <= rightArray[rightIndex] {
			result[arrayIndex] = leftArray[leftIndex]
			leftIndex++
		} else {
			result[arrayIndex] = rightArray[rightIndex]
			rightIndex++
		}
		arrayIndex++
	}

	for leftIndex < len(leftArray) {
		result[arrayIndex] = leftArray[leftIndex]
		leftIndex++
		arrayIndex++
	}

	for rightIndex < len(rightArray) {
		result[arrayIndex] = rightArray[rightIndex]
		rightIndex++
		arrayIndex++
	}

	return result
}
