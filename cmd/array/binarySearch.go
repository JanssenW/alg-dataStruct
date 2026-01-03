package main

import "fmt"

func main() {

	arr := []int{1, 3, 5, 6, 8, 9, 11, 15, 18, 20}
	target := 6
	search := binarySearch(arr, target)
	fmt.Println("Target", target, "found at index:", search)

}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + (high - low)) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = arr[mid] + 1

		} else if arr[mid] > target {
			high = arr[mid] - 1
		}

	}
	return -1

}
