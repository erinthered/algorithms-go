package merge

// Sort slice of ints using merge sort, O(nlogn) runtime
func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	first := Sort(arr[:len(arr)/2])
	second := Sort(arr[len(arr)/2:])

	return merge(first, second)
}

func merge(first, second []int) []int {
	// Make an array of the total size of the two slices
	merged := make([]int, len(first)+len(second))

	firstIndex, secondIndex, mergedIndex := 0, 0, 0
	// While we haven't finished with either array, check if element from first array is smaller than element from second array
	// If it is, add it to the merged array and continue, otherwise add the element from the second array
	// Note that this is not an in place sort if there are duplicate elements
	for firstIndex < len(first) && secondIndex < len(second) {
		if first[firstIndex] <= second[secondIndex] {
			merged[mergedIndex] = first[firstIndex]
			firstIndex++
		} else {
			merged[mergedIndex] = second[secondIndex]
			secondIndex++
		}
		mergedIndex++
	}

	// Check if there are still unmerged elements in first
	for i := firstIndex; i < len(first); i++ {
		merged[mergedIndex] = first[i]
		mergedIndex++
	}

	// Check if there are still unmerged elements in first
	for i := secondIndex; i < len(second); i++ {
		merged[mergedIndex] = second[i]
		mergedIndex++
	}

	return merged
}
