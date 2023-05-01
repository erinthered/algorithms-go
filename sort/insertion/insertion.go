package insertion

import "algorithms-go/sort/utility"

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
