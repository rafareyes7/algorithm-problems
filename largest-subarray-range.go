package main

import (
	"fmt"
	"reflect"
)

/*
  Write a function that takes in an array of integers and returns an array of
  length 2 representing the largest range of integers contained in that array.

  The first number in the output array should be the first number in the range,
  while the second number should be the last number in the range.


  A range of numbers is defined as a set of numbers that come right after each
  other in the set of real integers. For instance, the output array [2,6]
  represents the range {2, 3, 4, 5, 6}, which is a range of length 5.
  Note that numbers don't need to be sorted or adjacent in the input array in
  order to form a range.

  You can assume that there will only be one largest range.

  example:

  input: [1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6]
  output: [0, 7]

*/

// Time: O(n) | Space: O(n)
func largestSubarrayRange(arr []int) []int {
	result := []int{}
	largestRange := 0

	// Push all the elements into the map
	// with the num as the key and the value initialized as true, meaning
	// that the element has not been visited.
	numsMap := map[int]bool{}
	for _, num := range arr {
		numsMap[num] = true
	}

	// Start iterating the array checking if the current number is in the map
	// available to be visited (value = true), if so then start the range checking
	// by expanding <- ->
	for _, num := range arr {
		// If the number is not contained in the map, then continue.
		if !numsMap[num] {
			continue
		}
		// Since the current number is contained in the map and we are visiting it,
		// we inmediatly mark it as false
		numsMap[num] = false
		// Since the num is in the map, then the range is already of length 1
		currentRangeLength := 1
		// Expand to the left, checking if the previous values of the current num
		// are contained in the array and are able to visit (value = true)
		leftIdx := num - 1
		// <- while the number is in our map
		for numsMap[leftIdx] {
			// because we're visiting, must mark it as visited (false)
			numsMap[leftIdx] = false
			currentRangeLength++
			leftIdx--
		}
		// Expand to the right, checking if the next values of the current num
		// are contained in the array and are able to visit (value = true)
		rightIdx := num + 1
		// <- while the number is in our map
		for numsMap[rightIdx] {
			// because we're visiting, must mark it as visited (false)
			numsMap[rightIdx] = false
			currentRangeLength++
			rightIdx++
		}

		// Check the largest path
		if currentRangeLength > largestRange {
			// Is the largest range
			largestRange = currentRangeLength
			// Get the indexes of the start and end of the range
			result = []int{leftIdx + 1, rightIdx - 1}
		}
	}

	return result
}

func main() {

	want := []int{0, 7}
	got := largestSubarrayRange([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = []int{1, 1}
	got = largestSubarrayRange([]int{1})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = []int{1, 2}
	got = largestSubarrayRange([]int{1, 2})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = []int{-8, 19}
	got = largestSubarrayRange([]int{-7, -7, -7, -7, 8, -8, 0, 9, 19, -1, -3, 18, 17, 2, 10, 3, 12, 5, 16, 4, 11, -6, 8, 7, 6, 15, 12, 12, -5, 2, 1, 6, 13, 14, -4, -2})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

}
