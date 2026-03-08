package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left = 0; left <= right; {
		mid := (left + right) / 2

		// Check if target is present at the mid
		if arr[mid] == target {
			return mid
		}

		// If target is greater, ignore left half
		if arr[mid] < target {
			left = mid + 1
		}

		// If target is smaller, ignore right half
		else {
			right = mid - 1
		}
	}

	// If we reach here, then the element was not present
	return -1
}

func main() {
	// Test array
	arr := []int{2, 3, 4, 10, 40}
	target := 10

	// Function call
	r := binarySearch(arr, target)

	if r != -1 {
		fmt.Printf("Element is present at index %d\n", r)
	} else {
		fmt.Println("Element is not present in array")
	}
}