package main

import (
	"fmt"
	"reflect"
	"sort"
)

/* Write a function that takes in a non-empty array of integers and an integer representing
a target sum.

The function should find all triplets in the array that sum up to the target, and return a
two-dimensional array of all these triplets. The numbers in each triplet should be ordened in
ascending order, and the triplets themselves should be ordered in ascending order with respect to
the numbers they hold.

If no three numbers sum up to the target sum, the function should return an empty array.

example:

	input: [12,3,1,2,-6,5,-8,6], 0
	output: [[-8,2,6][-8,3,5][-6,1,5]]

*/

func threeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	triplets := [][]int{}
	for first := 0; first < len(array)-2; first++ {

		if first > 0 && array[first-1] == array[first] {
			// duplicated elem found -> [2, first:3, first+1:3, 5, ...]
			continue
		}

		mid, left := first+1, len(array)-1
		for mid < left {

			if mid > first+1 && array[mid-1] == array[mid] {
				// duplicated elem found -> [2, 3, 3, first:5, first+1:6, mid-1:7, mid:7, ...]
				mid++
				continue
			}

			currentSum := array[first] + array[mid] + array[left]
			if currentSum < target {
				mid++
			} else if currentSum > target {
				left--
			} else {
				triplets = append(triplets, []int{array[first], array[mid], array[left]})
				mid++
				left--
			}
		}
	}
	return triplets
}

func main() {

	want := [][]int{}
	got := threeNumberSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 5)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = [][]int{[]int{1, 2, 3}}
	got = threeNumberSum([]int{1, 2, 3}, 6)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = [][]int{[]int{-8, 2, 6}, []int{-8, 3, 5}, []int{-6, 1, 5}}
	got = threeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = [][]int{[]int{-2, 0, 2}}
	got = threeNumberSum([]int{-2, 0, 0, 2, 2}, 0)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = [][]int{[]int{0, 0, 0}}
	got = threeNumberSum([]int{0, 0, 0}, 0)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

}
