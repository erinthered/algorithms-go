package heap

import (
	"algorithms-go/sort/utility"
)

// Sort uses a heap to sort the elements of an array
func Sort(arr []int) []int {
	// build a max heap from array
	buildMaxHeap(arr)

	for i := len(arr) - 1; i > 0; i-- {
		utility.Swap(arr, 0, i) // move current root (largest element) to end of arr
		heapify(arr, i, 0)      // heapify arr[:i]
	}
	return arr
}

// heapify a tree from arr
func heapify(arr []int, size int, root int) {
	largest := root         // intitialize largest index to root
	left := (root * 2) + 1  // left child of root
	right := (root * 2) + 2 // right child of root

	// check if left child is larger than root
	if left < size && arr[largest] < arr[left] {
		largest = left
	}
	// check if right child is bigger than current largest
	if right < size && arr[largest] < arr[right] {
		largest = right
	}
	// if largest isn't still root, swap root and largest
	if largest != root {
		utility.Swap(arr, largest, root)

		// heapify subtree
		heapify(arr, size, largest)
	}

}

// buildMaxHeap builds a max heap from a given array
func buildMaxHeap(arr []int) {
	// get last non-leaf node
	lastNonLeafIndex := (len(arr) / 2) - 1

	// heapify non-leaf nodes in reverse level-order
	for i := lastNonLeafIndex; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
}
