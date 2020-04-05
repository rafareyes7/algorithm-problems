package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
Given an array of integers and an integer. Write a function that moves all instances of that
integer in the array to the end of the array.

The function should perfomr this in place (it should mutate the input array) and does not need
to mantain the order of the other integres.

example:
	input:  [2, 1, 2, 2, 2, 3, 4, 2], 2
	output: [1, 3, 4, 2, 2, 2, 2, 2]
*/

// Time: O(n) | Space: O(1)
func moveElementToEnd(array []int, toMove int) []int {
	firstIdx, lastIdx := 0, len(array)-1
	for firstIdx < lastIdx {
		if array[firstIdx] != toMove {
			firstIdx++
		}

		if array[lastIdx] == toMove {
			lastIdx--
		}

		//swap
		array[firstIdx], array[lastIdx] = array[lastIdx], array[firstIdx]
	}
	return array
}

func main() {

	want := []int{4, 3, 1, 2, 2, 2, 2, 2}
	got := moveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = []int{3, 3, 3, 3, 3, 3}
	got = moveElementToEnd([]int{3, 3, 3, 3, 3, 3}, 3)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = []int{1, 2, 4, 5, 3}
	got = moveElementToEnd([]int{3, 1, 2, 4, 5}, 3)
	sort.Ints(got[0:4])
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

}
