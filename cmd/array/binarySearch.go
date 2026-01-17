package main

import "fmt"

func main() {

	arr := []int{1, 3, 5, 6, 8, 9, 11, 15, 18, 20}
	target := 10
	search := binarySearch(arr, target, 0, len(arr)-1)
	// Essentialy binarySearch just need arr and target
	// However, to retrofit exponentinalBinarySearch, we need define low and high.
	fmt.Println("BinarySearch - Target", target, "found at index:", search)
	search = exponationalSearch(arr, target)
	fmt.Println("ExponentialSearch - Target", target, "found at index:", search)

}

func exponationalSearch(arr []int, target int) int {
	if arr[0] == target {
		return 0
	}

	low, high := 1, 1

	// Find range (subarray) for binary search by repeated doubling
	for high < len(arr) && arr[high] < target {
		low = high
		high = high * 2

	}

	if arr[high] == target {
		return high
	}

	high = min(high, len(arr)-1)
	// Pass a slice of the array to binary search, the min function is used to avoid index out of range error
	return binarySearch(arr, target, low, high)
}

func binarySearch(arr []int, target int, low int, high int) int {

	// In general, default to low should be 0 and high should be len(arr)-1.
	// We are not defing it as default parameters to retrofit exponational search.

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1

		} else if arr[mid] > target {
			high = mid - 1
		}

	}
	return -1

}
