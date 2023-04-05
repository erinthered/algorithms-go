package binary

import "fmt"

// BinarySearchRecursive returns true if the desired value is in the given data array or false if it is not
func binarySearchRecursive(data []int, value int) bool {
	if len(data) == 0 {
		return false
	}

	mid := len(data) / 2
	if data[mid] == value {
		return true
	} else if data[mid] < value && mid+1 < len(data) {
		return binarySearchRecursive(data[mid+1:], value)
	} else if data[mid] > value && mid > 0 {
		return binarySearchRecursive(data[:mid], value)
	}

	return false
}

func binarySearchIterative(data []int, value int) bool {
	var mid int
	low := 0
	high := len(data) - 1

	for low <= high {
		mid = (low + high) / 2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

func RunBinarySearch() {
	inputs := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{},
		{1, 8, 9, 10, 50, 108},
		{1},
		{5, 6, 7, 8, 9, 10},
		{101, 201, 301, 401},
		{1, 2, 3},
	}
	values := []int{5, 5, 10, 1, 7, 501, 3}
	expected := []bool{true, false, true, true, true, false, true}

	for i, input := range inputs {
		actual := binarySearchRecursive(input, values[i])
		var success string
		if actual == expected[i] {
			success = "success"
		} else {
			success = "fail"
		}
		fmt.Printf("Recursive Binary Search:: %s :: input: %v, desired value: %v, expected value: %t, actual value: %t\n", success, input, values[i], expected[i], actual)
	}

	for i, input := range inputs {
		actual := binarySearchIterative(input, values[i])
		var success string
		if actual == expected[i] {
			success = "success"
		} else {
			success = "fail"
		}
		fmt.Printf("Iterative Binary Search:: %s :: input: %v, desired value: %v, expected value: %t, actual value: %t\n", success, input, values[i], expected[i], actual)
	}
}
