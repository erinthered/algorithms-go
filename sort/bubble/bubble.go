package bubble

import "fmt"

type bubbleSort func([]int) []int

func bubbleSortIterative(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func bubbleSortRecursive(arr []int) []int {
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
	bubbleSortRecursive(arr[:len(arr)-1])

	return arr
}

// RunBubbleSort checks each element of an array with the next element and swaps them if next is smaller
// Runs in quadratic time (O(n^2))
func RunBubbleSort() {
	inputs := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{23, 18, 501, 112, 2, 1, 300},
	}
	expected := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 18, 23, 112, 300, 501},
	}
	testBubbleSort := func(bs bubbleSort) {
		for index, input := range inputs {
			fmt.Printf("Testing input: %v\n", input)
			output := bs(input)
			sorted := true
			for i, v := range output {
				if v != expected[index][i] {
					sorted = false
					break
				}
			}
			fmt.Printf("Sorted correctly: %t\t:: Output: %v\n\n", sorted, input)
		}
	}

	testBubbleSort(bubbleSortIterative)
	testBubbleSort(bubbleSortRecursive)
}
