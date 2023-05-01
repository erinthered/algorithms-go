package insertion

import "algorithms-go/sort/utility"

// Sort is insertion sort, average and worst case O(n^2), best case (sorted array) O(n)
func Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				utility.Swap(arr, j-1, j)
			}
		}
	}
	return arr
}
