package main

import "fmt"

// In general, it has bad performance, O(n^2) time complexity, and is not suitable for large datasets.
// However, it is simple to implement and understand, making it useful for small datasets or educational purposes.

// Time Complexity
// - Best Case: O(n) - occurs when the array is already sorted, and only one pass is needed to confirm this.
// - Woerst Case: O(n^2) - occurs when the array is sorted in reverse order, requiring maximum swaps.

// Space complexity is O(1) because it only requires a constant amount of additional space for the temporary variable used during swapping.

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArr := BubbleSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

func BubbleSort(arr []int) []int {
	size := len(arr)

	// Traverse through all array elements
	for i := 0; i < size; i++ {
		// Last i elements are already in place, so we decrease the range of comparison, hence size-i-1.
		for j := 0; j < size-i-1; j++ {
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}

	return arr

}
