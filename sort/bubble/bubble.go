package bubble

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if (j+1) < len(arr) && arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
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
	var sorted bool
	for index, input := range inputs {
		sorted = true
		fmt.Printf("Input: %v\t::", input)
		bubbleSort(input)
		for i, v := range input {
			if v != expected[index][i] {
				sorted = false
				break
			}
		}
		fmt.Printf("Sorted correctly: %t\t:: Output: %v\n", sorted, input)
	}
}
