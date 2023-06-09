package main

import (
	"algorithms-go/sort/bubble"
	"algorithms-go/sort/heap"
	"algorithms-go/sort/insertion"
	"algorithms-go/sort/merge"
	"algorithms-go/sort/quick"
	"algorithms-go/sort/selection"
	"fmt"
)

type sortFunction func([]int) []int

func main() {
	fmt.Println("Testing insertion sort")
	testSort(insertion.Sort)

	fmt.Println("Testing bubble sort")
	testSort(bubble.SortIterative)
	testSort(bubble.SortRecursive)
	testSort(bubble.SortRecursive)

	fmt.Println("Testing selection sort")
	testSort(selection.Sort)

	fmt.Println("Testing merge sort")
	testSort(merge.Sort)

	fmt.Println("Testing heap sort")
	testSort(heap.Sort)

	fmt.Println("Testing quick sort")
	testSort(quick.Sort)
}

func testSort(fun sortFunction) {
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
	for index, input := range inputs {
		fmt.Printf("Testing input %d: %v\n", index, input)
		output := fun(input)
		sorted := true
		for i, v := range output {
			if v != expected[index][i] {
				sorted = false
				break
			}
		}
		if !sorted {
			fmt.Printf("Did not sort correctly:: Output: %v\n\n", output)
		}
	}
}
