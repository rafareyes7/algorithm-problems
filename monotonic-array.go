package main

import (
	"fmt"
	"reflect"
)

/*
Write a function that takes in an array of integers and returns a boolean
representing whether the array is monotonic.

An array is said to be monotonic if its elements, from left to right,
are entirely non-increasing or entirely non-decreasing

example:
	input: [-1, -5, -10, -1100, -1101, -1102, -9001]
	output: true

	input: [1, 2, 0]
	output: false

*/

// Time: O(n) | Space: O(1)
func isMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	trendFound := false
	isIncreasing := false
	i := 1
	for i < len(array) {
		if array[i-1] == array[i] {
			i++
			continue
		} else if !trendFound {
			isIncreasing = array[i-1] < array[i]
			trendFound = true
		}

		if array[i-1] < array[i] != isIncreasing {
			return false
		}
		i++
	}

	return true
}

func main() {

	want := true
	got := isMonotonic([]int{-1, -5, -10, -1100, -1101, -1102, -9001})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = true
	got = isMonotonic([]int{1})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = true
	got = isMonotonic([]int{-1, -1, -1, -1, -1, -1, -1, -1})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = false
	got = isMonotonic([]int{1, 2, 0})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = false
	got = isMonotonic([]int{1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 7, 9, 10, 11})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}
}
