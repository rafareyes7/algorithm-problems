package main

import (
	"fmt"
	"reflect"
)

/*
  Write a function that takes in an array of integers and returns the length of
  the longest peak in the array.
  A peak is defined as adjacent integers in the array that are strictly
  increasing until they reach a tip (the highest value in the peak), at which
  point they become strictly decreasing. At least three integers are required to
  form a peak.

  For example, the integers 1, 4, 10, 2  form a peak, but the
  integers 4, 0, 10  don't and neither do the integers 1, 2, 2, 0.
  Similarly, the integers 1, 2, 3 don't form a peak because there aren't any
  strictly decreasing integers after the 3.

  input: [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
  output: 6 // 0, 10, 6, 5, -1, -3

*/

func longestPeak(arr []int) int {
	i, currentPeaklength, longestPeakLength := 1, 0, 0
	for i < len(arr)-1 {
		// Find peak
		isPeak := arr[i-1] < arr[i] && arr[i] > arr[i+1]
		if !isPeak {
			i++
			continue
		}

		// Expand to get the current length
		leftIdx := i - 2
		for leftIdx >= 0 && arr[leftIdx] < arr[leftIdx+1] { // increasing
			leftIdx--
		}
		rightIdx := i + 2
		for rightIdx < len(arr) && arr[rightIdx] < arr[rightIdx-1] { //decreaising
			rightIdx++
		}

		// Compute peak length (index diff)
		currentPeaklength = rightIdx - leftIdx - 1

		// Track the max peak length
		if currentPeaklength > longestPeakLength {
			longestPeakLength = currentPeaklength
		}

		// Keep iterating forward
		i = rightIdx

	}

	return longestPeakLength
}

func main() {
	want := len([]int{0, 10, 6, 5, -1, -3})
	got := longestPeak([]int{1, 2, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 0
	got = longestPeak([]int{})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 0
	got = longestPeak([]int{1, 2, 3, 4, 5, 6, 10, 100, 1000})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}
}
