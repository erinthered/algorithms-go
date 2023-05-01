package selection

import "algorithms-go/sort/utility"

// Sort input slice using selection sort, runs in O(n^2)
func Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			utility.Swap(arr, j, j-1)
		}
	}
	return arr
}
