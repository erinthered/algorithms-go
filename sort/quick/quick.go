package quick

import (
	"algorithms-go/sort/utility"
)

func Sort(arr []int) []int {
	// trivially sorted, empty array or single element
	if len(arr) <= 1 {
		return arr
	}
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// pi is partitioning index
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// choosing the last element of index as pivot
	pivot := arr[high]
	// index of smaller element
	i := low - 1

	for j := low; j <= high-1; j++ {
		// if current element is smaller than pivot, swap and increment smaller element
		if arr[j] < pivot {
			i++
			utility.Swap(arr, i, j)
		}
	}
	utility.Swap(arr, i+1, high)
	return i + 1
}
