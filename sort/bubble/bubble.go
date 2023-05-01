package bubble

// SortIterative is basic bubble sort, bubble up each value in the array
// Best, average, worst case all O(n^2)
func SortIterative(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// SortRecursive is bubble sort written recursively, same runtime as SortIterative
func SortRecursive(arr []int) []int {
	// Cover trivial and base cases: length of array is 0 or 1
	if len(arr) <= 1 {
		return arr
	}
	// If next element is smaller than current element, swap them
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	// Recursive call to bubblesort
	SortRecursive(arr[:len(arr)-1])

	return arr
}

// SortModified is a bubble sort that stops when array is sorted
// Improves best case to O(n), but worst and average remain O(n^n)
func SortModified(arr []int) []int {
	sorted := false
	for i := 0; i < len(arr); i++ {
		if sorted {
			break
		}
		sorted = true
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = false
			}
		}
	}
	return arr
}
