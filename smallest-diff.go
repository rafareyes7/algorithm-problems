package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

/*
	Write a funcrion that takes tow non-empty arrays of integers, find the pair of numbers
	(one from each array) whose absolute difference is closest to zero, and returns an array
	returning containing these two numbers, with the number from the first array in the first
	position

	example:

	input: A[-1, 5, 10, 20, 28, 3], B[26,134,135,15,17]
	output: [28, 26]
*/

func smallestDifference(arrayOne, arrayTwo []int) []int {
	sort.Ints(arrayOne)
	sort.Ints(arrayTwo)
	smallestPair := []int{}
	smallest, current := math.MaxInt32, math.MaxInt32
	idxOne, idxTwo := 0, 0

	for idxOne < len(arrayOne) && idxTwo < len(arrayTwo) {
		first, second := arrayOne[idxOne], arrayTwo[idxTwo]
		if first > second {
			current = first - second
			idxTwo++
		} else if second > first {
			current = second - first
			idxOne++
		} else {
			return []int{first, second}
		}

		if current < smallest {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}

func main() {

	want := []int{28, 26}
	got := smallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = []int{20, 21}
	got = smallestDifference([]int{0, 20}, []int{21, -2})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}

	want = []int{530, 530}
	got = smallestDifference([]int{10, 1000, 9124, 2142, 59, 24, 596, 591, 124, -123, 530}, []int{-1441, -124, -25, 1014, 1500, 660, 410, 245, 530})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v", want, got)
	}
}
