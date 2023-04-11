package bubble

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
