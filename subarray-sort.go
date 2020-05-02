package main

/*
  Write a function that takes in an array of at least two integers and that
  returns an array of the starting and ending indices of the smallest subarray
  in the input array that needs to be sorted in place in order for the entire
  input array to be sorted.

  If the input array is already sorted, the function should return [-1, -1]

  input: [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19]
  output: [3, 9]

*/

import (
	"fmt"
	"math"
	"reflect"
)

func subarraySort(arr []int) []int {
	minOutOfOrder, maxOutOfOrder := math.MaxInt32, math.MinInt32

	// Find the out of order nums
	for i, num := range arr {
		if isOutOfOrder(i, arr) {
			// Find min and max number of the out of order values
			minOutOfOrder = min(minOutOfOrder, num)
			maxOutOfOrder = max(maxOutOfOrder, num)
		}
	}

	// Edge case: the input array is already sorted
	// when minOutOfOrder remains the same initial value
	if minOutOfOrder == math.MaxInt32 {
		return []int{-1, -1}
	}

	// Since we've the min out of order value, we are gointo to find
	// the sorted position of that min at the final sorted array.
	// We can infer that all the values behind our min, are sorted and
	// are lower, so when we find a num greater than
	// our min then we have found the sorted position of our min.
	leftIdx := 0
	for arr[leftIdx] < minOutOfOrder {
		leftIdx++
	}

	// Since we've the max out of order value, we are gointo to find
	// the sorted position of that max at the final sorted array.
	// We can infer that all the values after our max, are sorted and
	// greater, so when we find a num lower than
	// our max then we have found the sorted position of our max.
	rightIdx := len(arr) - 1
	for arr[rightIdx] > maxOutOfOrder {
		rightIdx--
	}

	return []int{leftIdx, rightIdx}

}

func isOutOfOrder(currentIdx int, arr []int) bool {
	if currentIdx == 0 {
		return arr[currentIdx] > arr[currentIdx+1]
	} else if currentIdx == len(arr)-1 {
		return arr[currentIdx] < arr[currentIdx-1]
	}

	return arr[currentIdx] < arr[currentIdx-1] || arr[currentIdx] > arr[currentIdx+1]
}

func min(value1, value2 int) int {
	if value1 < value2 {
		return value1
	}

	return value2
}

func max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}

	return value2
}

func main() {

	want := []int{3, 9}
	got := subarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = []int{-1, -1}
	got = subarraySort([]int{1, 2, 3, 4, 5, 6, 7})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = []int{-1, -1}
	got = subarraySort([]int{})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

}
