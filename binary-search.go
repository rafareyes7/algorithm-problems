package main

import "fmt"

// binarySearch receives a sorted array and a taget number, if the target
// number is found within then array then the function returns its index
// in the array, otherwise the function returns nil
func binarySearch(arr []int, target int) *int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess < target {
			low = mid + 1
		} else if guess > target {
			high = mid - 1
		} else {
			return &mid
		}
	}
	return nil
}

func main() {
	want := 1
	got := binarySearch([]int{1, 3, 5, 7, 9}, 3)
	if got == nil || *got != want {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	got = binarySearch([]int{1, 3, 5, 7, 9}, -1)
	if got != nil {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

}
